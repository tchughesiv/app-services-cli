package decision

import (
	"testing"
)

func TestValidateName(t *testing.T) {

	type args struct {
		name string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Should throw error when exceeds 32 characters",
			args: args{
				name: "verylongdecisionnamef8d9dkf9dkc11dfs",
			},
			wantErr: true,
		},
		{
			name: "Should be valid when name is exactly 32 characters",
			args: args{
				name: "verylongdecisionnamef8d9dkf9dkc1",
			},
			wantErr: false,
		},
		{
			name: "Should be invalid when name is exactly 0 characters",
			args: args{
				name: "",
			},
			wantErr: true,
		},
		{
			name: "Should be invalid when using hyphens",
			args: args{
				name: "my-decision-instance",
			},
			wantErr: false,
		},
		{
			name: "Should be invalid when starts with number",
			args: args{
				name: "1my-decision-instance",
			},
			wantErr: true,
		},
		{
			name: "Should be invalid when includes uppercase letter",
			args: args{
				name: "my-Decision-instance",
			},
			wantErr: true,
		},
		{
			name: "Should be invalid when includes a space",
			args: args{
				name: "my-decision instance",
			},
			wantErr: true,
		},
		{
			name: "Should be invalid when includes a '.'",
			args: args{
				name: "my-decision.instance",
			},
			wantErr: true,
		},
		{
			name: "Should be invalid when starts with a '-'",
			args: args{
				name: "-my-decision-instance",
			},
			wantErr: true,
		},
		{
			name: "Should be invalid when ends with a '-'",
			args: args{
				name: "my-decision-instance-",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// nolint
			if err := ValidateName(tt.args.name); (err != nil) != tt.wantErr {
				t.Errorf("ValidateName() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

/*
func TestTransformRequest(t *testing.T) {
	hostWithSSLPort := "my-decision-url:443"
	hostWithNoPort := "my-decision-url"

	type args struct {
		decisionInstance *decisclient.DecisionRequest
	}
	tests := []struct {
		name                    string
		args                    args
		wantBootstrapServerHost string
	}{
		{
			name: "bootstrapServerHost should be transformed to empty string when nil",
			args: args{
				decisionInstance: &decisclient.DecisionRequest{
					BootstrapServerHost: nil,
				},
			},
			wantBootstrapServerHost: "",
		},
		{
			name: "bootstrapServerHost should be the same when SSL port already exists",
			args: args{
				decisionInstance: &decisclient.DecisionRequest{
					BootstrapServerHost: &hostWithSSLPort,
				},
			},
			wantBootstrapServerHost: hostWithSSLPort,
		},
		{
			name: "bootstrapServerHost should get SSL port when none exists",
			args: args{
				decisionInstance: &decisclient.DecisionRequest{
					BootstrapServerHost: &hostWithNoPort,
				},
			},
			wantBootstrapServerHost: hostWithSSLPort,
		},
	}
	for _, tt := range tests {
		// nolint
		t.Run(tt.name, func(t *testing.T) {
			transformedInstance := TransformDecisionRequest(tt.args.decisionInstance)

			if transformedInstance == nil {
				t.Errorf("Expected DecisionRequest type, but got nil")
			}

			transformedHost := transformedInstance.GetBootstrapServerHost()
			if tt.wantBootstrapServerHost != transformedHost {
				t.Errorf("Expected bootstrapServerHost: %v, got %v", tt.wantBootstrapServerHost, transformedHost)
			}
		})
	}
}

func TestTransformDecisionRequestListItems(t *testing.T) {
	hostWithSSLPort := "my-decision-url:443"
	hostWithNoPort := "my-decision-url"
	emptyHost := ""

	type args struct {
		items []decisclient.DecisionRequest
	}
	tests := []struct {
		name string
		args args
		want []decisclient.DecisionRequest
	}{
		{
			name: "Should return empty list when no decisions",
			args: args{
				items: []decisclient.DecisionRequest{},
			},
			want: []decisclient.DecisionRequest{},
		},
		{
			name: "Should update all bootstrapServerHost ports",
			args: args{
				items: []decisclient.DecisionRequest{
					{
						BootstrapServerHost: &emptyHost,
					},
					{
						BootstrapServerHost: &hostWithNoPort,
					},
					{
						BootstrapServerHost: &hostWithSSLPort,
					},
				},
			},
			want: []decisclient.DecisionRequest{
				{
					BootstrapServerHost: &emptyHost,
				},
				{
					BootstrapServerHost: &hostWithSSLPort,
				},
				{
					BootstrapServerHost: &hostWithSSLPort,
				},
			},
		},
	}
	for _, tt := range tests {
		// nolint
		t.Run(tt.name, func(t *testing.T) {
			gotItems := TransformDecisionRequestListItems(tt.args.items)

			if len(gotItems) != len(tt.want) {
				t.Fatalf("Expected number of items to be %v, got %v", len(gotItems), len(tt.want))
				return
			}

			for j, wantDecision := range tt.want {
				gotDecision := gotItems[j]

				gotBootstrapHost := gotDecision.GetBootstrapServerHost()
				wantBootstrapHost := wantDecision.GetBootstrapServerHost()

				if gotBootstrapHost != wantBootstrapHost {
					t.Fatalf("Expected BootstrapServerHost = %v, got %v", wantBootstrapHost, gotBootstrapHost)
				}
			}
		})
	}
}
*/
