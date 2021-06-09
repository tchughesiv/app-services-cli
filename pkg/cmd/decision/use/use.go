package use

import (
	"context"
	"errors"
	"fmt"

	decisclient "github.com/redhat-developer/app-services-cli/pkg/api/decis/client"
	"github.com/redhat-developer/app-services-cli/pkg/connection"
	"github.com/redhat-developer/app-services-cli/pkg/iostreams"
	"github.com/redhat-developer/app-services-cli/pkg/localize"

	"github.com/redhat-developer/app-services-cli/pkg/cmdutil"

	"github.com/redhat-developer/app-services-cli/pkg/decision"

	"github.com/spf13/cobra"

	"github.com/redhat-developer/app-services-cli/internal/config"
	"github.com/redhat-developer/app-services-cli/pkg/cmd/factory"
	"github.com/redhat-developer/app-services-cli/pkg/logging"
)

type Options struct {
	id          string
	name        string
	interactive bool

	IO         *iostreams.IOStreams
	Config     config.IConfig
	Connection factory.ConnectionFunc
	Logger     func() (logging.Logger, error)
	localizer  localize.Localizer
}

func NewUseCommand(f *factory.Factory) *cobra.Command {
	opts := &Options{
		Config:     f.Config,
		Connection: f.Connection,
		Logger:     f.Logger,
		IO:         f.IOStreams,
		localizer:  f.Localizer,
	}

	cmd := &cobra.Command{
		Use:     opts.localizer.MustLocalize("decision.use.cmd.use"),
		Short:   opts.localizer.MustLocalize("decision.use.cmd.shortDescription"),
		Long:    opts.localizer.MustLocalize("decision.use.cmd.longDescription"),
		Example: opts.localizer.MustLocalize("decision.use.cmd.example"),
		Args:    cobra.RangeArgs(0, 1),
		/*
			ValidArgsFunction: func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
				return cmdutil.FilterValidDecisions(f, toComplete)
			},
		*/
		ValidArgsFunction: func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
			return cmdutil.FilterValidDecisions(f)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) > 0 {
				opts.name = args[0]
			} else if opts.id == "" {
				if !opts.IO.CanPrompt() {
					return errors.New(opts.localizer.MustLocalize("decision.use.error.idOrNameRequired"))
				}
				opts.interactive = true
			}

			if opts.name != "" && opts.id != "" {
				return errors.New(opts.localizer.MustLocalize("decision.common.error.idAndNameCannotBeUsed"))
			}

			return runUse(opts)
		},
	}

	cmd.Flags().StringVar(&opts.id, "id", "", opts.localizer.MustLocalize("decision.use.flag.id"))

	return cmd
}

func runUse(opts *Options) error {

	if opts.interactive {
		// run the use command interactively
		err := runInteractivePrompt(opts)
		if err != nil {
			return err
		}
		// no Decision was selected, exit program
		if opts.name == "" {
			return nil
		}
	}

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

	var res *decisclient.DecisionRequest
	ctx := context.Background()
	if opts.name != "" {
		res, _, err = decision.GetDecisionByName(ctx, api.Decision(), opts.name)
		if err != nil {
			return err
		}
	} else {
		res, _, err = decision.GetDecisionByID(ctx, api.Decision(), opts.id)
		if err != nil {
			return err
		}
	}

	// build Decision config object from the response
	var decisionConfig config.DecisionConfig = config.DecisionConfig{
		ClusterID: res.GetId(),
	}

	nameTmplEntry := localize.NewEntry("Name", res.GetName())
	cfg.Services.Decision = &decisionConfig
	if err := opts.Config.Save(cfg); err != nil {
		saveErrMsg := opts.localizer.MustLocalize("decision.use.error.saveError", nameTmplEntry)
		return fmt.Errorf("%v: %w", saveErrMsg, err)
	}

	logger.Info(opts.localizer.MustLocalize("decision.use.log.info.useSuccess", nameTmplEntry))

	return nil
}

func runInteractivePrompt(opts *Options) error {
	logger, err := opts.Logger()
	if err != nil {
		return err
	}

	connection, err := opts.Connection(connection.DefaultConfigSkipMasAuth)
	if err != nil {
		return err
	}

	logger.Debug(opts.localizer.MustLocalize("common.log.debug.startingInteractivePrompt"))

	selectedDecision, err := decision.InteractiveSelect(connection, logger)
	if err != nil {
		return err
	}

	opts.name = selectedDecision.GetName()

	return nil
}
