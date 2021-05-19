package decision

import (
	"context"
	"fmt"

	"github.com/AlecAivazis/survey/v2"
	decisclient "github.com/redhat-developer/app-services-cli/pkg/api/decis/client"
	"github.com/redhat-developer/app-services-cli/pkg/connection"
	"github.com/redhat-developer/app-services-cli/pkg/logging"
)

/*
const (
	queryLimit = "1000"
)
*/

func InteractiveSelect(connection connection.Connection, logger logging.Logger) (*decisclient.DecisionRequest, error) {
	api := connection.API()

	response, _, err := api.Decision().ListDecisions(context.Background()).Execute()

	if err != nil {
		return nil, fmt.Errorf("unable to list Decision instances: %w", err)
	}

	/*
		if response.Size == nil {
			logger.Info("No Decision instances were found.")
			return nil, nil
		}
	*/

	decisions := []string{}
	for index := 0; index < len(response.Items); index++ {
		decisions = append(decisions, *response.Items[index].Name)
	}

	prompt := &survey.Select{
		Message:  "Select Decision instance to connect:",
		Options:  decisions,
		PageSize: 10,
	}

	var selectedDecisionIndex int
	err = survey.AskOne(prompt, &selectedDecisionIndex)
	if err != nil {
		return nil, err
	}

	selectedDecision := response.Items[selectedDecisionIndex]

	return &selectedDecision, nil
}
