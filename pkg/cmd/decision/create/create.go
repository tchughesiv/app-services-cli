package create

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"

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
	name        string
	descrip     string
	dmnFile     string
	kafkaSource string
	kafkaSink   string
	configs     map[string]string
	tags        map[string]string

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
)

var kind = "Decision"

// NewCreateCommand creates a new command for creating decisions.
func NewCreateCommand(f *factory.Factory) *cobra.Command {
	opts := &Options{
		IO:         f.IOStreams,
		Config:     f.Config,
		Connection: f.Connection,
		Logger:     f.Logger,
		localizer:  f.Localizer,
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
				if opts.dmnFile != "" {
					if err := decision.ValidateDMNfile(opts.dmnFile); err != nil {
						return err
					}
				}
			}

			if !opts.IO.CanPrompt() && opts.name == "" && opts.dmnFile == "" {
				return errors.New(opts.localizer.MustLocalize("decision.create.argument.name.error.requiredWhenNonInteractive"))
			} else if opts.name == "" || opts.descrip == "" || opts.dmnFile == "" {
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
	cmd.Flags().StringVar(&opts.dmnFile, flags.FlagDMN, "", opts.localizer.MustLocalize("decision.create.flag.dmn.description"))
	cmd.Flags().StringVar(&opts.kafkaSource, flags.FlagSource, "", opts.localizer.MustLocalize("decision.create.flag.kafkasource.description"))
	cmd.Flags().StringVar(&opts.kafkaSink, flags.FlagSink, "", opts.localizer.MustLocalize("decision.create.flag.kafkasink.description"))
	cmd.Flags().StringToStringVar(&opts.configs, flags.FlagConfig, map[string]string{}, opts.localizer.MustLocalize("decision.create.flag.config.description"))
	cmd.Flags().StringToStringVar(&opts.tags, flags.FlagTag, map[string]string{}, opts.localizer.MustLocalize("decision.create.flag.tag.description"))
	cmd.Flags().StringVarP(&opts.outputFormat, "output", "o", "json", opts.localizer.MustLocalize("decision.common.flag.output.description"))
	cmd.Flags().BoolVar(&opts.autoUse, "use", true, opts.localizer.MustLocalize("decision.create.flag.autoUse.description"))
	cmd.Flags().SortFlags = false

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

		payload = &decisclient.DecisionRequestPayload{
			Name:        opts.name,
			Kind:        &kind,
			Description: opts.descrip,
		}
		setEventing(opts.kafkaSource, opts.kafkaSink, payload)
		if err = setDMN(opts.dmnFile, payload); err != nil {
			return err
		}
	}
	setConfigsTags(opts.configs, opts.tags, payload)

	logger.Info(opts.localizer.MustLocalize("decision.create.log.debug.creatingDecision", localize.NewEntry("Name", opts.name)))

	a := api.Decision().CreateDecision(context.Background())
	a = a.DecisionRequestPayload(*payload)
	// a = a.Async(true)
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
	// set type to store the answers from the prompt with defaults
	answers := struct {
		Name        string
		Description string
		DMNFile     string
		KafkaSource string
		KafkaSink   string
	}{}

	answers.Name = opts.name
	if answers.Name == "" {
		promptName := &survey.Input{
			Message: opts.localizer.MustLocalize("decision.create.input.name.message"),
			Help:    opts.localizer.MustLocalize("decision.create.input.name.help"),
		}
		err = survey.AskOne(promptName, &answers.Name, survey.WithValidator(pkgDecision.ValidateName))
		if err != nil {
			return nil, err
		}
	}

	answers.Description = opts.descrip
	if answers.Description == "" {
		descripPrompt := &survey.Input{
			Message: opts.localizer.MustLocalize("decision.create.input.descrip.message"),
			Default: defaultDescrip,
			//Options: cloudProviderNames,
		}
		err = survey.AskOne(descripPrompt, &answers.Description)
		if err != nil {
			return nil, err
		}
	}

	answers.DMNFile = opts.dmnFile
	if answers.DMNFile == "" {
		dmnPrompt := &survey.Input{
			Message: opts.localizer.MustLocalize("decision.create.input.dmn.message"),
			Help:    opts.localizer.MustLocalize("decision.create.input.dmn.help"),
		}
		err = survey.AskOne(dmnPrompt, &answers.DMNFile, survey.WithValidator(pkgDecision.ValidateDMNfile))
		if err != nil {
			return nil, err
		}
	}

	answers.KafkaSource = opts.kafkaSource
	if answers.KafkaSource == "" {
		kSourcePrompt := &survey.Input{Message: "Kafka Source:", Help: "optional"}
		err = survey.AskOne(kSourcePrompt, &answers.KafkaSource)
		if err != nil {
			return nil, err
		}
	}

	answers.KafkaSink = opts.kafkaSink
	if answers.KafkaSink == "" {
		kSinkPrompt := &survey.Input{Message: "Kafka Sink:", Help: "optional"}
		err = survey.AskOne(kSinkPrompt, &answers.KafkaSink)
		if err != nil {
			return nil, err
		}
	}

	payload = &decisclient.DecisionRequestPayload{
		Name:        answers.Name,
		Kind:        &kind,
		Description: answers.Description,
	}
	if err = setDMN(answers.DMNFile, payload); err != nil {
		return nil, err
	}
	setEventing(answers.KafkaSource, answers.KafkaSink, payload)

	return payload, nil
}

// setDMN ...
func setDMN(dmnFile string, payload *decisclient.DecisionRequestPayload) error {
	if dmnFile != "" {
		dmn, err := ioutil.ReadFile(dmnFile)
		if err != nil {
			return err
		}
		payload.Model = decisclient.DecisionRequestPayloadAllOfModel{
			Dmn: string(dmn),
		}
	}
	return nil
}

// setEventing ...
func setEventing(source, sink string, payload *decisclient.DecisionRequestPayload) {
	if source != "" || sink != "" {
		payload.Eventing = &decisclient.DecisionRequestAllOfEventing{
			Kafka: &decisclient.DecisionRequestAllOfEventingKafka{},
		}
		if source != "" {
			payload.Eventing.Kafka.Source = &source
		}
		if sink != "" {
			payload.Eventing.Kafka.Sink = &sink
		}
	}
}

// setConfigsTags ...
func setConfigsTags(configs, tags map[string]string, payload *decisclient.DecisionRequestPayload) {
	if configs != nil {
		payload.Configuration = &configs
	}
	if tags != nil {
		payload.Tags = &tags
	}
}
