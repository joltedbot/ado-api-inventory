package main

import (
	"bufio"
	"crypto/tls"
	"encoding/json"
	"fmt"
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

	resp, _ := secureClient.Do(req)
	if err != nil {
		return "", "", err
	}

	defer resp.Body.Close()

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

func writeToFile(fileName string, data string) {
	file, err := os.Create(OUTPUT_DIRECTORY + "/" + fileName)
	if err != nil {
		println(err)
	}

	defer file.Close()

	_, err = file.WriteString(data)
	if err != nil {
		println(err)
	}
}

func getEndpointStruct[T any](endpoint EndPoint, results APIResults[T], authentication string) (APIResults[T], error) {

	continuationToken := ""

	for {
		loopResult := APIResults[T]{}
		response, token, err := apiCall(endpoint.resource, apiURL(endpoint.isGraph, endpoint.urlBase, endpoint.resource, endpoint.parameters), continuationToken, authentication)
		if err != nil {
			return APIResults[T]{}, err
		}

		println(response)

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
