package main

import (
	"crypto/tls"
	"net/http"
	"os"
	"testing"
)

func TestApiURL(t *testing.T) {
	tests := []struct {
		baseUrl         string
		organizationUrl string
		endpoint        string
		parameters      string
		expectedPrefix  string
		expectError     bool
	}{
		{"https://dev.azure.com/", "org", "projects", "", "https://dev.azure.com/org/_apis/projects?api-version=7.2-preview", false},
		{"https://vssps.dev.azure.com/", "org", "graph/users", "top=10", "https://vssps.dev.azure.com/org/_apis/graph/users?api-version=7.2-preview&top=10", false},
		{"http://[::1]:namedport", "org", "projects", "", "", true},
	}

	for _, tt := range tests {
		got, err := apiURL(tt.baseUrl, tt.organizationUrl, tt.endpoint, tt.parameters)
		if tt.expectError {
			if err == nil {
				t.Errorf("apiURL() expected error for %v, but got none", tt.baseUrl)
			}
			continue
		}

		if err != nil {
			t.Errorf("apiURL() unexpected error: %v", err)
			continue
		}

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

	transport, ok := client.Transport.(*http.Transport)
	if !ok {
		t.Fatal("Expected *http.Transport, got different type")
	}

	if transport.TLSClientConfig.MinVersion != tls.VersionTLS12 {
		t.Errorf("Expected TLS min version %v, got %v", tls.VersionTLS12, transport.TLSClientConfig.MinVersion)
	}

	if transport.MaxIdleConns != 10 {
		t.Errorf("Expected MaxIdleConns 10, got %v", transport.MaxIdleConns)
	}
}

func TestWriteToFileAppend(t *testing.T) {
	testFile := "testfile_append.txt"
	testData1 := "first line\n"
	testData2 := "second line\n"

	err := os.Mkdir(OUTPUT_DIRECTORY, 0755)
	if err != nil && !os.IsExist(err) {
		t.Fatalf("Failed to create output directory: %v", err)
	}
	defer os.Remove(OUTPUT_DIRECTORY + "/" + testFile)

	// Write initial data
	writeToFile(testFile, testData1, false)

	// Append additional data
	writeToFile(testFile, testData2, true)

	// Read and verify combined content
	f, err := os.Open(OUTPUT_DIRECTORY + "/" + testFile)
	if err != nil {
		t.Fatalf("Failed to open test file: %v", err)
	}
	defer deferCloseFile(f)

	buf := make([]byte, len(testData1)+len(testData2))
	_, err = f.Read(buf)
	if err != nil {
		t.Fatalf("File read error: %v", err)
	}

	expected := testData1 + testData2
	if string(buf) != expected {
		t.Errorf("writeToFile append contents = %v, want %v", string(buf), expected)
	}
}

func TestApiURLEdgeCases(t *testing.T) {
	tests := []struct {
		name            string
		baseUrl         string
		organizationUrl string
		endpoint        string
		parameters      string
		expected        string
		expectError     bool
	}{
		{
			name:            "empty parameters",
			baseUrl:         "https://dev.azure.com/",
			organizationUrl: "org",
			endpoint:        "projects",
			parameters:      "",
			expected:        "https://dev.azure.com/org/_apis/projects?api-version=7.2-preview",
			expectError:     false,
		},
		{
			name:            "base URL without trailing slash",
			baseUrl:         "https://dev.azure.com",
			organizationUrl: "org",
			endpoint:        "projects",
			parameters:      "",
			expected:        "https://dev.azure.com/org/_apis/projects?api-version=7.2-preview",
			expectError:     false,
		},
		{
			name:            "special characters in organization",
			baseUrl:         "https://dev.azure.com/",
			organizationUrl: "my-org_test",
			endpoint:        "projects",
			parameters:      "",
			expected:        "https://dev.azure.com/my-org_test/_apis/projects?api-version=7.2-preview",
			expectError:     false,
		},
		{
			name:            "multiple parameters",
			baseUrl:         "https://dev.azure.com/",
			organizationUrl: "org",
			endpoint:        "projects",
			parameters:      "top=10&skip=5",
			expected:        "https://dev.azure.com/org/_apis/projects?api-version=7.2-preview&top=10&skip=5",
			expectError:     false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := apiURL(tt.baseUrl, tt.organizationUrl, tt.endpoint, tt.parameters)
			if tt.expectError {
				if err == nil {
					t.Errorf("apiURL() expected error for %s, but got none", tt.name)
				}
				return
			}

			if err != nil {
				t.Errorf("apiURL() unexpected error for %s: %v", tt.name, err)
				return
			}

			if got != tt.expected {
				t.Errorf("apiURL() %s = %v, want %v", tt.name, got, tt.expected)
			}
		})
	}
}
