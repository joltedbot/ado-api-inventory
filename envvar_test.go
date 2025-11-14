package main

import (
	"os"
	"testing"
)

func TestGetAndValidateEnvVarsValid(t *testing.T) {
	// Use valid UUID v4 values (version nibble = 4)
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
