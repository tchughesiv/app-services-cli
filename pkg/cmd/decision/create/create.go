package create

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/redhat-developer/app-services-cli/pkg/localize"

	"github.com/redhat-developer/app-services-cli/pkg/cmd/decision/flags"
	"github.com/redhat-developer/app-services-cli/pkg/cmd/flag"
	flagutil "github.com/redhat-developer/app-services-cli/pkg/cmdutil/flags"
	"github.com/redhat-developer/app-services-cli/pkg/connection"
	"github.com/redhat-developer/app-services-cli/pkg/decision"

	decisclient "github.com/redhat-developer/app-services-cli/pkg/api/decis/client"

	"github.com/AlecAivazis/survey/v2"
	"github.com/redhat-developer/app-services-cli/pkg/dump"
	"github.com/redhat-developer/app-services-cli/pkg/iostreams"

	pkgDecision "github.com/redhat-developer/app-services-cli/pkg/decision"
	"github.com/redhat-developer/app-services-cli/pkg/logging"

	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"

	"github.com/redhat-developer/app-services-cli/internal/config"
	"github.com/redhat-developer/app-services-cli/pkg/cmd/factory"
	"github.com/redhat-developer/app-services-cli/pkg/cmdutil"
)

type Options struct {
	name    string
	descrip string
	dmn     string
	// multiAZ bool

	outputFormat string
	autoUse      bool

	interactive bool

	IO         *iostreams.IOStreams
	Config     config.IConfig
	Connection factory.ConnectionFunc
	Logger     func() (logging.Logger, error)
	localizer  localize.Localizer
}

const (
	// default Decision instance values
	defaultDescrip = "A test decision"
	defaultDMN     = "<xml></xml>"
	// defaultMultiAZ  = true
	// defaultRegion   = "us-east-1"
	// defaultProvider = "aws"
)

var (
	source = "test"
	sink   = "test"
	kind   = "Decision"
)

// NewCreateCommand creates a new command for creating decisions.
func NewCreateCommand(f *factory.Factory) *cobra.Command {
	opts := &Options{
		IO:         f.IOStreams,
		Config:     f.Config,
		Connection: f.Connection,
		Logger:     f.Logger,
		localizer:  f.Localizer,

		// multiAZ: defaultMultiAZ,
	}

	cmd := &cobra.Command{
		Use:     opts.localizer.MustLocalize("decision.create.cmd.use"),
		Short:   opts.localizer.MustLocalize("decision.create.cmd.shortDescription"),
		Long:    opts.localizer.MustLocalize("decision.create.cmd.longDescription"),
		Example: opts.localizer.MustLocalize("decision.create.cmd.example"),
		Args:    cobra.RangeArgs(0, 1),
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) > 0 {
				opts.name = args[0]

				if err := decision.ValidateName(opts.name); err != nil {
					return err
				}
			}

			if !opts.IO.CanPrompt() && opts.name == "" {
				return errors.New(opts.localizer.MustLocalize("decision.create.argument.name.error.requiredWhenNonInteractive"))
			} else if opts.name == "" && opts.descrip == "" && opts.dmn == "" {
				opts.interactive = true
			}

			validOutputFormats := flagutil.ValidOutputFormats
			if opts.outputFormat != "" && !flagutil.IsValidInput(opts.outputFormat, validOutputFormats...) {
				return flag.InvalidValueError("output", opts.outputFormat, validOutputFormats...)
			}

			return runCreate(opts)
		},
	}

	cmd.Flags().StringVar(&opts.descrip, flags.FlagDescrip, "", opts.localizer.MustLocalize("decision.create.flag.descrip.description"))
	cmd.Flags().StringVar(&opts.dmn, flags.FlagDMN, "", opts.localizer.MustLocalize("decision.create.flag.dmn.description"))
	cmd.Flags().StringVarP(&opts.outputFormat, "output", "o", "json", opts.localizer.MustLocalize("decision.common.flag.output.description"))
	cmd.Flags().BoolVar(&opts.autoUse, "use", true, opts.localizer.MustLocalize("decision.create.flag.autoUse.description"))

	return cmd
}

// nolint:funlen
func runCreate(opts *Options) error {
	logger, err := opts.Logger()
	if err != nil {
		return err
	}

	cfg, err := opts.Config.Load()
	if err != nil {
		return err
	}

	connection, err := opts.Connection(connection.DefaultConfigSkipMasAuth)
	if err != nil {
		return err
	}

	api := connection.API()

	/*
		// the user must have accepted the terms and conditions from the provider
		// before they can create a decision instance
		termsAccepted, termsURL, err := checkTermsAccepted(opts.Connection)
		if err != nil {
			return err
		}
		if !termsAccepted && termsURL != "" {
			logger.Info(opts.localizer.MustLocalize("decision.create.log.info.termsCheck", localize.NewEntry("TermsURL", termsURL)))
			return nil
		}
	*/
	var payload *decisclient.DecisionRequestPayload
	if opts.interactive {
		logger.Debug()

		payload, err = promptDecisionPayload(opts)
		if err != nil {
			return err
		}

	} else {
		if opts.descrip == "" {
			opts.descrip = opts.name
		}
		if opts.dmn == "" {
			opts.dmn = defaultDMN
		}

		payload = &decisclient.DecisionRequestPayload{
			Name:        opts.name,
			Kind:        &kind,
			Description: opts.descrip,
			Model:       decisclient.DecisionRequestPayloadAllOfModel{Dmn: opts.dmn},
			Eventing: &decisclient.DecisionRequestAllOfEventing{
				Kafka: &decisclient.DecisionRequestAllOfEventingKafka{
					Source: &source,
					Sink:   &sink,
				},
			},
			Configuration: &map[string]string{"test": "test"},
			Tags: &map[string]string{
				"test1": "test",
				"test2": "test",
			},
			// Region:        &opts.region,
			// CloudProvider: &opts.provider,
			// MultiAz:       &opts.multiAZ,
		}
	}

	logger.Info(opts.localizer.MustLocalize("decision.create.log.debug.creatingDecision", localize.NewEntry("Name", opts.name)))

	a := api.Decision().CreateDecision(context.Background())
	a = a.DecisionRequestPayload(*payload)
	a = a.Async(true)
	response, _, err := a.Execute()

	if err != nil {
		return err
	}

	logger.Info(opts.localizer.MustLocalize("decision.create.info.successMessage", localize.NewEntry("Name", response.GetName())))

	switch opts.outputFormat {
	case "json":
		data, _ := json.MarshalIndent(response, "", cmdutil.DefaultJSONIndent)
		_ = dump.JSON(opts.IO.Out, data)
	case "yaml", "yml":
		data, _ := yaml.Marshal(response)
		_ = dump.YAML(opts.IO.Out, data)
	}

	decisionCfg := &config.DecisionConfig{
		ClusterID: response.GetId(),
	}

	if opts.autoUse {
		logger.Debug(opts.localizer.MustLocalize("decision.create.debug.autoUseSetMessage"))
		cfg.Services.Decision = decisionCfg
		if err := opts.Config.Save(cfg); err != nil {
			return fmt.Errorf("%v: %w", opts.localizer.MustLocalize("decision.common.error.couldNotUseDecision"), err)
		}
	} else {
		logger.Debug(opts.localizer.MustLocalize("decision.create.debug.autoUseNotSetMessage"))
	}

	return nil
}

// Show a prompt to allow the user to interactively insert the data for their Decision
func promptDecisionPayload(opts *Options) (payload *decisclient.DecisionRequestPayload, err error) {
	// connection, err := opts.Connection(connection.DefaultConfigSkipMasAuth)
	// if err != nil {
	//	return nil, err
	// }

	// api := connection.API()

	// set type to store the answers from the prompt with defaults
	answers := struct {
		Name        string
		Description string
		DMN         string
		// Region        string
		// MultiAZ       bool
		// CloudProvider string
	}{
		Description: defaultDescrip,
		DMN:         defaultDMN,
	}

	promptName := &survey.Input{
		Message: opts.localizer.MustLocalize("decision.create.input.name.message"),
		Help:    opts.localizer.MustLocalize("decision.create.input.name.help"),
	}
	err = survey.AskOne(promptName, &answers.Name, survey.WithValidator(pkgDecision.ValidateName))
	if err != nil {
		return nil, err
	}

	// fetch all cloud available providers
	// cloudProviderResponse, _, err := api.Decision().ListCloudProviders(context.Background()).Execute()
	// if err != nil {
	//	return nil, err
	// }

	// cloudProviders := cloudProviderResponse.GetItems()
	// cloudProviderNames := cloudproviderutil.GetEnabledNames(cloudProviders)
	descripPrompt := &survey.Input{
		Message: opts.localizer.MustLocalize("decision.create.input.descrip.message"),
		//Options: cloudProviderNames,
	}
	err = survey.AskOne(descripPrompt, &answers.Description)
	if err != nil {
		return nil, err
	}

	// get the selected provider type from the name selected
	// selectedCloudProvider := cloudproviderutil.FindByName(cloudProviders, answers.CloudProvider)

	// nolint
	// cloudRegionResponse, _, err := api.Decision().ListCloudProviderRegions(context.Background(), selectedCloudProvider.GetId()).Execute()
	// if err != nil {
	//	return nil, err
	// }

	// regions := cloudRegionResponse.GetItems()
	// regionIDs := cloudregionutil.GetEnabledIDs(regions)
	dmnPrompt := &survey.Input{
		Message: opts.localizer.MustLocalize("decision.create.input.dmn.message"),
		// Options: regionIDs,
		Help: opts.localizer.MustLocalize("decision.create.input.dmn.help"),
	}
	err = survey.AskOne(dmnPrompt, &answers.DMN)
	if err != nil {
		return nil, err
	}

	payload = &decisclient.DecisionRequestPayload{
		Name:        answers.Name,
		Kind:        &kind,
		Description: answers.Description,
		Model: decisclient.DecisionRequestPayloadAllOfModel{
			Dmn: answers.DMN,
		},
		Eventing: &decisclient.DecisionRequestAllOfEventing{
			Kafka: &decisclient.DecisionRequestAllOfEventingKafka{
				Source: &source,
				Sink:   &sink,
			},
		},
		Configuration: &map[string]string{"test": "test"},
		Tags:          &map[string]string{"test": "test"},
		// Region:        &answers.Region,
		// CloudProvider: &answers.CloudProvider,
		// MultiAz:       &answers.MultiAZ,
	}

	return payload, nil
}

/*
func checkTermsAccepted(connFunc factory.ConnectionFunc) (accepted bool, redirectURI string, err error) {
	conn, err := connFunc(connection.DefaultConfigSkipMasAuth)
	if err != nil {
		return false, "", err
	}

	termsReview, _, err := conn.API().AccountMgmt().
		ApiAuthorizationsV1SelfTermsReviewPost(context.Background()).
		SelfTermsReview(amsclient.SelfTermsReview{
			EventCode: &build.TermsReviewEventCode,
			SiteCode:  &build.TermsReviewSiteCode,
		}).
		Execute()
	if err != nil {
		return false, "", err
	}

	if !termsReview.GetTermsAvailable() && !termsReview.GetTermsRequired() {
		return true, "", nil
	}

	if !termsReview.HasRedirectUrl() {
		return false, "", errors.New("terms must be signed, but there is no terms URL")
	}

	return false, termsReview.GetRedirectUrl(), nil
}
*/
