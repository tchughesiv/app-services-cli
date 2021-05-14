package list

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/redhat-developer/app-services-cli/internal/config"
	decisclient "github.com/redhat-developer/app-services-cli/pkg/api/decis/client"
	"github.com/redhat-developer/app-services-cli/pkg/cmd/factory"
	"github.com/redhat-developer/app-services-cli/pkg/cmd/flag"
	flagutil "github.com/redhat-developer/app-services-cli/pkg/cmdutil/flags"
	"github.com/redhat-developer/app-services-cli/pkg/connection"
	"github.com/redhat-developer/app-services-cli/pkg/decision"
	"github.com/redhat-developer/app-services-cli/pkg/dump"
	"github.com/redhat-developer/app-services-cli/pkg/iostreams"
	"github.com/redhat-developer/app-services-cli/pkg/localize"
	"github.com/redhat-developer/app-services-cli/pkg/logging"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"
)

// row is the details of a Decision instance needed to print to a table
type decisionRow struct {
	ID      string `json:"id" header:"ID"`
	Name    string `json:"name" header:"Name"`
	Version string `json:"version" header:"Version"`
	Status  string `json:"status" header:"Status"`
	// Owner         string `json:"owner" header:"Owner"`
	// CloudProvider string `json:"cloud_provider" header:"Cloud Provider"`
	// Region        string `json:"region" header:"Region"`
}

type options struct {
	outputFormat string
	page         int
	limit        int
	search       string

	IO         *iostreams.IOStreams
	Config     config.IConfig
	Connection factory.ConnectionFunc
	Logger     func() (logging.Logger, error)
	localizer  localize.Localizer
}

// NewListCommand creates a new command for listing decisions.
func NewListCommand(f *factory.Factory) *cobra.Command {
	opts := &options{
		page:       0,
		limit:      100,
		search:     "",
		Config:     f.Config,
		Connection: f.Connection,
		Logger:     f.Logger,
		IO:         f.IOStreams,
		localizer:  f.Localizer,
	}

	cmd := &cobra.Command{
		Use:   opts.localizer.MustLocalize("decision.list.cmd.use"),
		Short: opts.localizer.MustLocalize("decision.list.cmd.shortDescription"),
		Long:  opts.localizer.MustLocalize("decision.list.cmd.longDescription"),
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			if opts.outputFormat != "" && !flagutil.IsValidInput(opts.outputFormat, flagutil.ValidOutputFormats...) {
				return flag.InvalidValueError("output", opts.outputFormat, flagutil.ValidOutputFormats...)
			}

			if err := decision.ValidateSearchInput(opts.search); err != nil {
				return err
			}

			return runList(opts)
		},
	}

	cmd.Flags().StringVarP(&opts.outputFormat, "output", "o", "", opts.localizer.MustLocalize("decisions.common.flag.output.description"))
	cmd.Flags().IntVarP(&opts.page, "page", "", 0, opts.localizer.MustLocalize("decision.list.flag.page"))
	cmd.Flags().IntVarP(&opts.limit, "limit", "", 100, opts.localizer.MustLocalize("decision.list.flag.limit"))
	cmd.Flags().StringVarP(&opts.search, "search", "", "", opts.localizer.MustLocalize("decision.list.flag.search"))

	return cmd
}

func runList(opts *options) error {
	logger, err := opts.Logger()
	if err != nil {
		return err
	}
	logger.DebugEnabled()

	connection, err := opts.Connection(connection.DefaultConfigSkipMasAuth)
	if err != nil {
		return err
	}

	api := connection.API()

	a := api.Decision().ListDecisions(context.Background())
	a = a.Page(strconv.Itoa(opts.page))
	a = a.Size(strconv.Itoa(opts.limit))

	if opts.search != "" {
		query := buildQuery(opts.search)
		logger.Debug(opts.localizer.MustLocalize("decision.list.log.debug.filteringDecisionList", localize.NewEntry("Search", query)))
		a = a.Search(query)
	}

	response, _, err := a.Execute()
	if err != nil {
		return err
	}

	if response.Size == 0 && opts.outputFormat == "" {
		logger.Info(opts.localizer.MustLocalize("decision.common.log.info.noDecisionInstances"))
		return nil
	}

	switch opts.outputFormat {
	case "json":
		data, _ := json.Marshal(response)
		_ = dump.JSON(opts.IO.Out, data)
	case "yaml", "yml":
		data, _ := yaml.Marshal(response)
		_ = dump.YAML(opts.IO.Out, data)
	default:
		rows := mapResponseItemsToRows(response.GetItems())
		dump.Table(opts.IO.Out, rows)
		logger.Info("")
	}

	return nil
}

func mapResponseItemsToRows(decisions []decisclient.DecisionRequest) []decisionRow {
	rows := []decisionRow{}

	for _, d := range decisions {
		row := decisionRow{
			ID:      d.GetId(),
			Name:    d.GetName(),
			Version: d.GetVersion(),
			Status:  d.GetStatus(),
			//Owner:         d.GetOwner(),
			//CloudProvider: d.GetCloudProvider(),
			//Region:        d.GetRegion(),
		}

		rows = append(rows, row)
	}

	return rows
}

func buildQuery(search string) string {

	queryString := fmt.Sprintf(
		"name like %%%[1]v%% or owner like %%%[1]v%% or cloud_provider like %%%[1]v%% or region like %%%[1]v%% or status like %%%[1]v%%",
		search,
	)

	return queryString

}
