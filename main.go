package main

import (
	"fmt"
	"os"
	"sync"
)

const OUTPUT_DIRECTORY = "output"

func main() {

	environment, err := getAndValidateEnvVars()
	if err != nil {
		println("You must set the ADO_TENANT_ID, ADO_CLIENT_ID, ADO_CLIENT_SECRET, and ADO_ORGANIZATION environment variables correctly before running this program.")
		os.Exit(1)
	}

	err = os.Mkdir(OUTPUT_DIRECTORY, 0700)
	if err != nil && !os.IsExist(err) {
		panic(err)
	}

	adoToken := getADOToken(environment.TenantId, environment.ClientId, environment.ClientSecret)

	projectIDs, err := getProjects(environment.Organization, adoToken)

	if err != nil {
		panic("Could not retrieve Project data and can not continue processings. Exiting.")
	}

	fmt.Println("----------------------------------------")
	fmt.Println("       API Call Return Statuses         ")
	fmt.Println("----------------------------------------")

	var wg sync.WaitGroup

	wg.Add(9)

	go getUsers(environment.Organization, adoToken, &wg)
	go getGroups(environment.Organization, adoToken, &wg)
	go getTeams(environment.Organization, adoToken, &wg)
	go getWiki(environment.Organization, adoToken, &wg)
	go getArtifactFeeds(environment.Organization, adoToken, &wg)
	go getRepositories(environment.Organization, adoToken, projectIDs, &wg)
	go getPipelines(environment.Organization, adoToken, projectIDs, &wg)
	go getBoards(environment.Organization, adoToken, projectIDs, &wg)
	go getTestPlans(environment.Organization, adoToken, projectIDs, &wg)

	wg.Wait()
}
