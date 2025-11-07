package main

import (
	"os"
)

const OUTPUT_DIRECTORY = "output"

func main() {

	tenant := os.Getenv("ADO_TENANT_ID")
	clientId := os.Getenv("ADO_CLIENT_ID")
	clientSecret := os.Getenv("ADO_CLIENT_SECRET")
	organization := os.Getenv("ADO_ORGANIZATION")

	err := os.Mkdir(OUTPUT_DIRECTORY, 0755)
	if err != nil && !os.IsExist(err) {
		panic(err)
	}

	adoToken := getADOToken(tenant, clientId, clientSecret)

	go getUsers(organization, adoToken)
	go getProjects(organization, adoToken)
	go getTeams(organization, adoToken)
	go getRepositories(organization, adoToken)

}
