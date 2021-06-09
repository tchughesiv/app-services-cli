package decision

import (
	"errors"
	"os"
	"regexp"

	"github.com/redhat-developer/app-services-cli/pkg/common/commonerr"
	"github.com/redhat-developer/app-services-cli/pkg/decision/decisionerr"
)

var (
	validNameRegexp   = regexp.MustCompile(`^[a-z]([-a-z0-9]*[a-z0-9])?$`)
	validSearchRegexp = regexp.MustCompile(`^([a-zA-Z0-9-_%]*[a-zA-Z0-9-_%])?$`)
)

// ValidateName validates the proposed name of a Decision instance
func ValidateName(val interface{}) error {
	name, ok := val.(string)

	if !ok {
		return commonerr.NewCastError(val, "string")
	}

	if len(name) < 1 || len(name) > 32 {
		return errors.New("Decision instance name must be between 1 and 32 characters")
	}

	matched := validNameRegexp.MatchString(name)

	if matched {
		return nil
	}

	return decisionerr.InvalidNameError(name)
}

// ValidateDMNfile validates the proposed name of a Decision instance
func ValidateDMNfile(val interface{}) error {
	name, ok := val.(string)
	if !ok {
		return commonerr.NewCastError(val, "string")
	}
	if _, err := os.Stat(name); err != nil {
		if os.IsNotExist(err) {
			return decisionerr.FileNotFoundError(name)
		}
	}
	return nil
}

/*
// TransformDecisionRequestListItems modifies fields fields from a list of decision instances
// The main transformation is appending ":443" to the Bootstrap Server URL
func TransformDecisionRequestListItems(items []decisclient.DecisionRequest) []decisclient.DecisionRequest {
	for i := range items {
		decision := items[i]
		decision = *TransformDecisionRequest(&decision)
		items[i] = decision
	}

	return items
}

// TransformDecisionRequest modifies fields from the DecisionRequest payload object
// The main transformation is appending ":443" to the Bootstrap Server URL
func TransformDecisionRequest(decision *decisclient.DecisionRequest) *decisclient.DecisionRequest {
	bootstrapHost := decision.GetBootstrapServerHost()

	if bootstrapHost == "" {
		return decision
	}

	if !strings.HasSuffix(bootstrapHost, ":443") {
		hostURL := fmt.Sprintf("%v:443", bootstrapHost)
		decision.SetBootstrapServerHost(hostURL)
	}

	return decision
}
*/

// ValidateSearchInput validates the text provided to filter the Decision instances
func ValidateSearchInput(val interface{}) error {
	search, ok := val.(string)

	if !ok {
		return commonerr.NewCastError(val, "string")
	}

	matched := validSearchRegexp.MatchString(search)

	if matched {
		return nil
	}

	return decisionerr.InvalidSearchValueError(search)
}
