package main

import (
	"os"
	"testing"
)

func TestApiURL(t *testing.T) {
	tests := []struct {
		isGraph         bool
		organizationUrl string
		endpoint        string
		parameters      string
		expectedPrefix  string
	}{
		{false, "org", "projects", "", "https://dev.azure.com/org/_apis/projects?api-version=7.2-preview"},
		{true, "org", "graph/users", "top=10", "https://vssps.dev.azure.com/org/_apis/graph/users?api-version=7.2-preview&top=10"},
	}

	for _, tt := range tests {
		got := apiURL(tt.isGraph, tt.organizationUrl, tt.endpoint, tt.parameters)
		if got != tt.expectedPrefix {
			t.Errorf("apiURL() = %v, want %v", got, tt.expectedPrefix)
		}
	}
}

func TestWriteToFile(t *testing.T) {
	testFile := "testfile.txt"
	testData := "hello world"

	err := os.Mkdir(OUTPUT_DIRECTORY, 0755)
	if err != nil && !os.IsExist(err) {
		t.Fatalf("Failed to create output directory: %v", err)
	}

	writeToFile(testFile, testData, false)

	f, err := os.Open(OUTPUT_DIRECTORY + "/" + testFile)

	if err != nil {
		t.Fatalf("writeToFile did not create file: %v", err)
	}

	defer deferCloseFile(f)
	buf := make([]byte, len(testData))
	_, err = f.Read(buf)
	if err != nil {
		t.Fatalf("writeToFile file read error: %v", err)
	}
	if string(buf) != testData {
		t.Errorf("writeToFile contents = %v, want %v", string(buf), testData)
	}

	err = os.Remove(OUTPUT_DIRECTORY + "/" + testFile)
	if err != nil {
		t.Fatalf("Failed to remove test file, remove it manually: %v", err)
	}
}

func TestNewSecureHTTPClient(t *testing.T) {
	client := newSecureHTTPClient()
	if client.Timeout != 30*1e9 {
		t.Errorf("newSecureHTTPClient Timeout = %v, want %v", client.Timeout, 30*1e9)
	}
}
