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

	return base + organizationUrl + "/_apis/" + endpoint + "?api-version=7.2-preview" + parameters

}

func apiCall(name string, url string, authentication string) (string, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		panic(err)
	}

	req.Header.Add("Authorization", "Bearer "+authentication)

	secureClient := newSecureHTTPClient()

	resp, _ := secureClient.Do(req)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	fmt.Println(name+" status:", resp.Status)

	scanner := bufio.NewScanner(resp.Body)

	responseBody := ""

	for i := 0; scanner.Scan() && i < 5; i++ {
		responseBody += scanner.Text()
	}

	if err := scanner.Err(); err != nil {
		return "", err
	}

	return responseBody, nil
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

	response, err := apiCall(endpoint.resource, apiURL(endpoint.isGraph, endpoint.urlBase, endpoint.resource, ""), authentication)
	if err != nil {
		return APIResults[T]{}, err
	}

	err = json.Unmarshal([]byte(response), &results)
	if err != nil {
		return APIResults[T]{}, err
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
