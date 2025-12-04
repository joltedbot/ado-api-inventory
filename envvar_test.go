package main

import (
	"os"
	"testing"
)

func TestGetAndValidateEnvVarsValid(t *testing.T) {
	_ = os.Setenv("ADO_TENANT_ID", "3fa85f64-5717-4562-b3fc-2c963f66afa6")
	_ = os.Setenv("ADO_CLIENT_ID", "4b825dc6-9d54-4f54-9f5b-123456789abc")
	_ = os.Setenv("ADO_CLIENT_SECRET", "secretValue123~")
	_ = os.Setenv("ADO_ORGANIZATION", "org-name-123")

	_, err := getAndValidateEnvVars()
	if err != nil {
		t.Errorf("getAndValidateEnvVars() valid input returned error: %v", err)
	}
}

func TestGetAndValidateEnvVarsInvalid(t *testing.T) {
	defer suppressTestOutput()()

	_ = os.Setenv("ADO_TENANT_ID", "12345")
	_ = os.Setenv("ADO_CLIENT_ID", "12345")
	_ = os.Setenv("ADO_CLIENT_SECRET", "12345")
	_ = os.Setenv("ADO_ORGANIZATION", "12345")

	_, err := getAndValidateEnvVars()
	if err == nil {
		t.Error("getAndValidateEnvVars() invalid input did not return error")
	}
}

func TestGetAndValidateEnvVarsMissing(t *testing.T) {
	defer suppressTestOutput()()

	// Clear all environment variables
	os.Unsetenv("ADO_TENANT_ID")
	os.Unsetenv("ADO_CLIENT_ID")
	os.Unsetenv("ADO_CLIENT_SECRET")
	os.Unsetenv("ADO_ORGANIZATION")

	_, err := getAndValidateEnvVars()
	if err == nil {
		t.Error("getAndValidateEnvVars() missing env vars did not return error")
	}
}

func TestGetAndValidateEnvVarsBoundaryCases(t *testing.T) {
	defer suppressTestOutput()()

	tests := []struct {
		name         string
		tenantId     string
		clientId     string
		clientSecret string
		organization string
		expectError  bool
	}{
		{
			name:         "valid max length tenant ID",
			tenantId:     "3fa85f64-5717-4562-b3fc-2c963f66afa6",
			clientId:     "4b825dc6-9d54-4f54-9f5b-123456789abc",
			clientSecret: "secretValue123~",
			organization: "org-name-123",
			expectError:  false,
		},
		{
			name:         "valid max length organization",
			tenantId:     "3fa85f64-5717-4562-b3fc-2c963f66afa6",
			clientId:     "4b825dc6-9d54-4f54-9f5b-123456789abc",
			clientSecret: "secretValue123~",
			organization: "a1234567890123456789012345678901234567890",
			expectError:  false,
		},

		{
			name:         "invalid UUID format",
			tenantId:     "not-a-uuid",
			clientId:     "4b825dc6-9d54-4f54-9f5b-123456789abc",
			clientSecret: "secretValue123~",
			organization: "org-name-123",
			expectError:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			os.Setenv("ADO_TENANT_ID", tt.tenantId)
			os.Setenv("ADO_CLIENT_ID", tt.clientId)
			os.Setenv("ADO_CLIENT_SECRET", tt.clientSecret)
			os.Setenv("ADO_ORGANIZATION", tt.organization)

			_, err := getAndValidateEnvVars()
			if tt.expectError && err == nil {
				t.Errorf("getAndValidateEnvVars() %s expected error but got none", tt.name)
			}
			if !tt.expectError && err != nil {
				t.Errorf("getAndValidateEnvVars() %s unexpected error: %v", tt.name, err)
			}
		})
	}
}
