package delete

import (
	"context"
	"errors"
	"fmt"

	decisclient "github.com/redhat-developer/app-services-cli/pkg/api/decis/client"
	"github.com/redhat-developer/app-services-cli/pkg/cmdutil"
	"github.com/redhat-developer/app-services-cli/pkg/connection"
	"github.com/redhat-developer/app-services-cli/pkg/localize"

	"github.com/redhat-developer/app-services-cli/pkg/decision"
	"github.com/redhat-developer/app-services-cli/pkg/iostreams"

	"github.com/redhat-developer/app-services-cli/pkg/logging"

	"github.com/AlecAivazis/survey/v2"
	"github.com/redhat-developer/app-services-cli/internal/config"
	"github.com/redhat-developer/app-services-cli/pkg/cmd/factory"
	"github.com/redhat-developer/app-services-cli/pkg/cmd/flag"
	"github.com/spf13/cobra"
)

type options struct {
	id    string
	name  string
	force bool

	IO         *iostreams.IOStreams
	Config     config.IConfig
	Connection factory.ConnectionFunc
	Logger     func() (logging.Logger, error)
	localizer  localize.Localizer
}

// NewDeleteCommand command for deleting decisions.
func NewDeleteCommand(f *factory.Factory) *cobra.Command {
	opts := &options{
		Config:     f.Config,
		Connection: f.Connection,
		Logger:     f.Logger,
		IO:         f.IOStreams,
		localizer:  f.Localizer,
	}

	cmd := &cobra.Command{
		Use:     opts.localizer.MustLocalize("decision.delete.cmd.use"),
		Short:   opts.localizer.MustLocalize("decision.delete.cmd.shortDescription"),
		Long:    opts.localizer.MustLocalize("decision.delete.cmd.longDescription"),
		Example: opts.localizer.MustLocalize("decision.delete.cmd.example"),
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
			if !opts.IO.CanPrompt() && !opts.force {
				return flag.RequiredWhenNonInteractiveError("yes")
			}

			if len(args) > 0 {
				opts.name = args[0]
			}

			if opts.name != "" && opts.id != "" {
				return errors.New(opts.localizer.MustLocalize("decision.common.error.idAndNameCannotBeUsed"))
			}

			if opts.id != "" || opts.name != "" {
				return runDelete(opts)
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

			return runDelete(opts)
		},
	}

	cmd.Flags().StringVar(&opts.id, "id", "", opts.localizer.MustLocalize("decision.delete.flag.id"))
	cmd.Flags().BoolVarP(&opts.force, "yes", "y", false, opts.localizer.MustLocalize("decision.delete.flag.yes"))

	return cmd
}

func runDelete(opts *options) error {
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

	var response *decisclient.DecisionRequest
	ctx := context.Background()
	if opts.name != "" {
		response, _, err = decision.GetDecisionByName(ctx, api.Decision(), opts.name)
		if err != nil {
			return err
		}
	} else {
		response, _, err = decision.GetDecisionByID(ctx, api.Decision(), opts.id)
		if err != nil {
			return err
		}
	}

	decisionName := response.GetName()

	logger.Info(opts.localizer.MustLocalize("decision.delete.log.info.deleting", localize.NewEntry("Name", decisionName)))
	logger.Info("")

	if !opts.force {
		promptConfirmName := &survey.Input{
			Message: opts.localizer.MustLocalize("decision.delete.input.confirmName.message"),
		}

		var confirmedDecisionName string
		err = survey.AskOne(promptConfirmName, &confirmedDecisionName)
		if err != nil {
			return err
		}

		if confirmedDecisionName != decisionName {
			logger.Info(opts.localizer.MustLocalize("decision.delete.log.info.incorrectNameConfirmation"))
			return nil
		}
	}

	// delete the Decision
	logger.Debug(opts.localizer.MustLocalize("decision.delete.log.debug.deletingDecision"), fmt.Sprintf("\"%s\"", decisionName))
	a := api.Decision().DeleteDecisionById(context.Background(), response.GetId())
	a = a.Async(true)
	_, _, err = a.Execute()

	if err != nil {
		return err
	}

	logger.Info(opts.localizer.MustLocalize("decision.delete.log.info.deleteSuccess", localize.NewEntry("Name", decisionName)))

	currentDecision := cfg.Services.Decision
	// this is not the current cluster, our work here is done
	if currentDecision == nil || currentDecision.ClusterID != response.GetId() {
		return nil
	}

	// the Decision that was deleted is set as the user's current cluster
	// since it was deleted it should be removed from the config
	cfg.Services.Decision = nil
	err = opts.Config.Save(cfg)
	if err != nil {
		return err
	}

	return nil
}