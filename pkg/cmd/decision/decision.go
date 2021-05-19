// Package decision instance contains commands for interacting with cluster logic of the service directly instead of through the
// REST API exposed via the serve command.
package decision

import (
	"github.com/spf13/cobra"

	"github.com/redhat-developer/app-services-cli/pkg/cmd/decision/list"
	"github.com/redhat-developer/app-services-cli/pkg/cmd/decision/use"
	"github.com/redhat-developer/app-services-cli/pkg/cmd/factory"
)

func NewDecisionCommand(f *factory.Factory) *cobra.Command {
	cmd := &cobra.Command{
		Use:   f.Localizer.MustLocalize("decision.cmd.use"),
		Short: f.Localizer.MustLocalize("decision.cmd.shortDescription"),
		Args:  cobra.MinimumNArgs(1),
	}

	// add sub-commands
	cmd.AddCommand(
		//create.NewCreateCommand(f),
		//describe.NewDescribeCommand(f),
		//delete.NewDeleteCommand(f),
		list.NewListCommand(f),
		use.NewUseCommand(f),
	)

	return cmd
}
