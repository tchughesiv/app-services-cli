package decision

import (
	"context"
	"net/http"

	"github.com/redhat-developer/app-services-cli/pkg/api/decis"
	decisclient "github.com/redhat-developer/app-services-cli/pkg/api/decis/client"
	"github.com/redhat-developer/app-services-cli/pkg/decision/decisionerr"
)

func GetDecisionByID(ctx context.Context, api decisclient.DefaultApi, id string) (*decisclient.DecisionRequest, *http.Response, error) {
	r := api.GetDecisionById(ctx, id)

	decisionReq, httpResponse, err := r.Execute()
	if decis.IsErr(err, decis.ErrorNotFound) {
		return nil, httpResponse, decisionerr.NotFoundByIDError(id)
	}

	return &decisionReq, httpResponse, err
}

func GetDecisionByName(ctx context.Context, api decisclient.DefaultApi, name string) (*decisclient.DecisionRequest, *http.Response, error) {
	r := api.ListDecisions(ctx)
	// r = r.Search(fmt.Sprintf("name = %v", name))
	decisionList, httpResponse, err := r.Execute()
	if err != nil {
		return nil, httpResponse, err
	}

	/*
		if decisionList.GetTotal() == 0 {
			return nil, nil, decisionerr.NotFoundByNameError(name)
		}

		items := decisionList.GetItems()
		decisionReq := items[0]

		return &decisionReq, httpResponse, err
	*/
	for _, decision := range decisionList.GetItems() {
		if name == decision.GetName() {
			return &decision, httpResponse, err
		}
	}

	return nil, nil, decisionerr.NotFoundByNameError(name)
}
