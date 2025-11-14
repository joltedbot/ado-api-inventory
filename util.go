package main

import (
	"bufio"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

func apiURL(isGraph bool, organizationUrl string, endpoint string, parameters string) string {

	base := "https://dev.azure.com/"
	if isGraph {
		base = "https://vssps.dev.azure.com/"
	}

	if parameters != "" {
		parameters = "&" + parameters
	}

	return base + organizationUrl + "/_apis/" + endpoint + "?api-version=7.2-preview" + parameters

}

func apiCall(name string, url string, continuationToken string, authentication string) (string, string, error) {

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		panic(err)
	}

	req.Header.Add("Authorization", "Bearer "+authentication)

	if continuationToken != "" {
		req.Header.Add("x-ms-continuationtoken", continuationToken)
	}

	secureClient := newSecureHTTPClient()

	resp, err := secureClient.Do(req)
	if err != nil {
		return "", "", err
	}

	defer deferCloseResponseBody(resp.Body)

	fmt.Println(name+" status:", resp.Status)

	continuationHeader := resp.Header.Get("x-ms-continuationtoken")

	responseBody := ""
	scanner := bufio.NewScanner(resp.Body)

	for scanner.Scan() {
		responseBody += scanner.Text()
	}

	if err := scanner.Err(); err != nil {
		return "", "", err
	}

	return responseBody, continuationHeader, nil
}

func writeToFile(fileName string, data string, append bool) error {

	filePath := OUTPUT_DIRECTORY + "/" + fileName
	var file *os.File
	var err error

	if append {
		file, err = os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0640)
	} else {
		file, err = os.Create(filePath)
	}

	if err != nil {
		return err
	}

	defer deferCloseFile(file)

	_, err = file.WriteString(data)
	if err != nil {
		return err
	}

	err = file.Sync()

	return err
}

func getEndpointStruct[T any](endpoint EndPoint, results APIResults[T], authentication string) (APIResults[T], error) {

	continuationToken := ""

	for {
		loopResult := APIResults[T]{}

		response, token, err := apiCall(endpoint.resource, apiURL(endpoint.isGraph, endpoint.urlBase, endpoint.resource, endpoint.parameters), continuationToken, authentication)
		if err != nil {
			return APIResults[T]{}, err
		}

		continuationToken = token

		err = json.Unmarshal([]byte(response), &loopResult)
		if err != nil {
			return APIResults[T]{}, err
		}

		results.Count += loopResult.Count
		results.Value = append(results.Value, loopResult.Value...)

		if continuationToken == "" {
			break
		}

	}
	return results, nil

}

func fetchAndExport[T any](
	endpoint EndPoint,
	authentication string,
	iteration int,
	formatRow func(T) string,
) error {
	result := APIResults[T]{Value: []T{}}

	result, err := getEndpointStruct(endpoint, result, authentication)
	if err != nil {
		log.Printf("getEndpointStruct(): %s - %s\n", endpoint.resource, err)
		return err
	}

	var output string
	fileAppend := true

	if iteration == 0 {
		output = endpoint.headerRow + "\n"
		fileAppend = false
	}

	for _, item := range result.Value {
		output += formatRow(item)
	}

	err = writeToFile(endpoint.fileName, output, fileAppend)
	return err
}

func newSecureHTTPClient() *http.Client {
	return &http.Client{
		Timeout: 30 * time.Second,
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				MinVersion: tls.VersionTLS12,
			},
			MaxIdleConns:    10,
			IdleConnTimeout: 120 * time.Second,
		},
	}
}

func deferCloseFile(file *os.File) {
	err := file.Close()
	if err != nil {
		log.Println(err)
	}
}

func deferCloseResponseBody(body io.ReadCloser) {
	err := body.Close()
	if err != nil {
		log.Println(err)
	}
}

func suppressTestOutput() func() {
	null, _ := os.Open(os.DevNull)
	sout := os.Stdout
	serr := os.Stderr
	os.Stdout = null
	os.Stderr = null
	log.SetOutput(null)
	return func() {
		defer null.Close()
		os.Stdout = sout
		os.Stderr = serr
		log.SetOutput(os.Stderr)
	}
}
