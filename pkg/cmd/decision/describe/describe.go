package describe

import (
	"context"
	"encoding/json"
	"errors"

	flagutil "github.com/redhat-developer/app-services-cli/pkg/cmdutil/flags"
	"github.com/redhat-developer/app-services-cli/pkg/connection"
	"github.com/redhat-developer/app-services-cli/pkg/iostreams"
	"github.com/redhat-developer/app-services-cli/pkg/localize"

	"github.com/redhat-developer/app-services-cli/pkg/cmd/flag"

	"github.com/redhat-developer/app-services-cli/internal/config"
	decisclient "github.com/redhat-developer/app-services-cli/pkg/api/decis/client"
	"github.com/redhat-developer/app-services-cli/pkg/cmd/factory"
	"github.com/redhat-developer/app-services-cli/pkg/cmdutil"
	"github.com/redhat-developer/app-services-cli/pkg/decision"
	"github.com/redhat-developer/app-services-cli/pkg/dump"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"
)

type Options struct {
	id           string
	name         string
	outputFormat string

	IO         *iostreams.IOStreams
	Config     config.IConfig
	Connection factory.ConnectionFunc
	localizer  localize.Localizer
}

// NewDescribeCommand describes a Decision instance, either by passing an `--id flag`
// or by using the decision instance set in the config, if any
func NewDescribeCommand(f *factory.Factory) *cobra.Command {
	opts := &Options{
		Config:     f.Config,
		Connection: f.Connection,
		IO:         f.IOStreams,
		localizer:  f.Localizer,
	}

	cmd := &cobra.Command{
		Use:     opts.localizer.MustLocalize("decision.describe.cmd.use"),
		Short:   opts.localizer.MustLocalize("decision.describe.cmd.shortDescription"),
		Long:    opts.localizer.MustLocalize("decision.describe.cmd.longDescription"),
		Example: opts.localizer.MustLocalize("decision.describe.cmd.example"),
		Args:    cobra.RangeArgs(0, 1),
		// Dynamic completion of the Decision name
		/*
			ValidArgsFunction: func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
				return cmdutil.FilterValidDecisions(f, toComplete)
			},
		*/
		ValidArgsFunction: func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
			return cmdutil.FilterValidDecisions(f)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			validOutputFormats := flagutil.ValidOutputFormats
			if opts.outputFormat != "" && !flagutil.IsValidInput(opts.outputFormat, validOutputFormats...) {
				return flag.InvalidValueError("output", opts.outputFormat, validOutputFormats...)
			}

			if len(args) > 0 {
				opts.name = args[0]
			}

			if opts.name != "" && opts.id != "" {
				return errors.New(opts.localizer.MustLocalize("decision.common.error.idAndNameCannotBeUsed"))
			}

			if opts.id != "" || opts.name != "" {
				return runDescribe(opts)
			}

			cfg, err := opts.Config.Load()
			if err != nil {
				return err
			}

			var decisionConfig *config.DecisionConfig
			if cfg.Services.Decision == decisionConfig || cfg.Services.Decision.ClusterID == "" {
				return errors.New(opts.localizer.MustLocalize("decision.common.error.noDecisionSelected"))
			}

			opts.id = cfg.Services.Decision.ClusterID

			return runDescribe(opts)
		},
	}

	cmd.Flags().StringVarP(&opts.outputFormat, "output", "o", "json", opts.localizer.MustLocalize("decision.common.flag.output.description"))
	cmd.Flags().StringVar(&opts.id, "id", "", opts.localizer.MustLocalize("decision.describe.flag.id"))

	return cmd
}

func runDescribe(opts *Options) error {
	connection, err := opts.Connection(connection.DefaultConfigSkipMasAuth)
	if err != nil {
		return err
	}

	api := connection.API()

	var decisionInstance *decisclient.DecisionRequest
	ctx := context.Background()
	if opts.name != "" {
		decisionInstance, _, err = decision.GetDecisionByName(ctx, api.Decision(), opts.name)
		if err != nil {
			return err
		}
	} else {
		decisionInstance, _, err = decision.GetDecisionByID(ctx, api.Decision(), opts.id)
		if err != nil {
			return err
		}
	}

	return printDecision(decisionInstance, opts)
}

func printDecision(decision *decisclient.DecisionRequest, opts *Options) error {
	switch opts.outputFormat {
	case "yaml", "yml":
		data, err := yaml.Marshal(decision)
		if err != nil {
			return err
		}
		return dump.YAML(opts.IO.Out, data)
	default:
		data, err := json.Marshal(decision)
		if err != nil {
			return err
		}
		return dump.JSON(opts.IO.Out, data)
	}
}
