package main

import (
	"fmt"
	"log"
	"sync"
)

func getProjects(organizationUrl string, authentication string) []string {

	projectList := APIResults[project]{
		Value: []project{},
	}

	endpoint := EndPoint{
		resource:   "projects",
		parameters: "",
		urlBase:    organizationUrl,
		isGraph:    false,
	}

	projectList, err := getEndpointStruct(endpoint, projectList, authentication)
	if err != nil {
		log.Printf("getEndpointStruct(): %s - %s\n", endpoint.resource, err)
	}

	fileName := "projects.csv"
	output := "Id,Name,Description,State,Visibility,LastUpdate,URL\n"
	var projectIDs []string

	for _, project := range projectList.Value {
		projectIDs = append(projectIDs, project.Id)
		output += fmt.Sprintf("%s,%s,%s,%s,%s,%s,%s\n", project.Id, project.Name, project.Description, project.State, project.Visibility, project.LastUpdate, project.URL)
	}

	writeToFile(fileName, output)

	return projectIDs
}

func getTeams(organizationUrl string, authentication string, wg *sync.WaitGroup) {
	defer wg.Done()

	teamsList := APIResults[teams]{
		Value: []teams{},
	}

	endpoint := EndPoint{
		resource:   "teams",
		parameters: "",
		urlBase:    organizationUrl,
		isGraph:    false,
	}

	teamsList, err := getEndpointStruct(endpoint, teamsList, authentication)
	if err != nil {
		log.Printf("getEndpointStruct(): %s - %s\n", endpoint.resource, err)
	}

	teamsFileName := "teams.csv"
	teamsOutput := "Id,Name,Description,Project ID,Project Name,URL,Identity Id,Identity URL\n"

	for _, team := range teamsList.Value {
		teamsOutput += fmt.Sprintf("%s,%s,%s,%s,%s,%s,%s,%s\n", team.Id, team.Name, team.Description, team.ProjectId, team.ProjectName, team.URL, team.Id, team.IdentityUrl)
	}

	writeToFile(teamsFileName, teamsOutput)
}

func getUsers(organizationUrl string, authentication string, wg *sync.WaitGroup) {
	defer wg.Done()

	userList := APIResults[users]{
		Value: []users{},
	}

	endpoint := EndPoint{
		resource:   "graph/users",
		parameters: "",
		urlBase:    organizationUrl,
		isGraph:    true,
	}

	userList, err := getEndpointStruct(endpoint, userList, authentication)
	if err != nil {
		log.Printf("getEndpointStruct(): %s - %s\n", endpoint.resource, err)
	}

	fileName := "graph-users.csv"
	output := "Display Name, Descriptor, Description, Principal Name, Mail Address, Subject Kind, Domain, Origin, Origin ID, URL\n"

	for _, user := range userList.Value {
		output += fmt.Sprintf("%s,%s,%s,%s,%s,%s,%s,%s,%s,%s\n", user.DisplayName, user.Descriptor, user.Description, user.PrincipalName, user.MailAddress, user.SubjectKind, user.Domain, user.Origin, user.OriginId, user.URL)
	}

	writeToFile(fileName, output)

}

func getGroups(organizationUrl string, authentication string, wg *sync.WaitGroup) {
	defer wg.Done()

	groupList := APIResults[groups]{
		Value: []groups{},
	}

	endpoint := EndPoint{
		resource:   "graph/groups",
		parameters: "",
		urlBase:    organizationUrl,
		isGraph:    true,
	}

	groupList, err := getEndpointStruct(endpoint, groupList, authentication)
	if err != nil {
		log.Printf("getEndpointStruct(): %s - %s\n", endpoint.resource, err)
	}

	fileName := "graph-groups.csv"
	output := "Display Name, Descriptor, Description, Principal Name, Mail Address, Subject Kind, Domain, Origin, Origin ID, URL\n"

	for _, group := range groupList.Value {
		output += fmt.Sprintf("%s,%s,%s,%s,%s,%s,%s,%s,%s,%s\n", group.DisplayName, group.Descriptor, group.Description, group.PrincipalName, group.MailAddress, group.SubjectKind, group.Domain, group.Origin, group.OriginId, group.URL)
	}

	writeToFile(fileName, output)

}

func getPipelines(organizationUrl string, authentication string, projectIDs []string, wg *sync.WaitGroup) {
	defer wg.Done()

	pipelineList := APIResults[pipelines]{
		Value: []pipelines{},
	}

	fileName := "pipelines.csv"
	output := "ID, Name, Folder, Revision, Project ID, URL, Configuration Type\n"

	for _, projectID := range projectIDs {

		endpoint := EndPoint{
			resource:   "pipelines",
			parameters: "",
			urlBase:    organizationUrl + "/" + projectID,
			isGraph:    false,
		}

		pipelineList, err := getEndpointStruct(endpoint, pipelineList, authentication)
		if err != nil {
			log.Printf("getEndpointStruct(): %s - %s\n", endpoint.resource, err)
		}

		for _, pipeline := range pipelineList.Value {
			output += fmt.Sprintf("%d,%s,%s,%d,%s,%s,%s\n", pipeline.Id, pipeline.Name, pipeline.Folder, pipeline.Revision, projectID, pipeline.URL, pipeline.Configuration.Type)
		}

	}

	writeToFile(fileName, output)

}

func getRepositories(organizationUrl string, authentication string, projectIDs []string, wg *sync.WaitGroup) {
	defer wg.Done()

	repositoryList := APIResults[repository]{
		Value: []repository{},
	}

	fileName := "repositories.csv"
	output := "Id,Name,Project ID,Created Date,Size,Default Branch,URL,Remote URL,SSH URL,Valid Remote URLs,Web URL,Is Disabled,Is Fork,Is In Maintenance,Parent Repository ID,Project ID\n"

	for _, projectID := range projectIDs {
		endpoint := EndPoint{
			resource:   "git/repositories",
			parameters: "",
			urlBase:    organizationUrl + "/" + projectID,
			isGraph:    false,
		}

		repositoryList, err := getEndpointStruct(endpoint, repositoryList, authentication)
		if err != nil {
			log.Printf("getEndpointStruct(): %s - %s\n", endpoint.resource, err)
		}

		for _, repository := range repositoryList.Value {
			output += fmt.Sprintf("%s,%s,%s,%s,%d,%s,%s,%s,%s,%s,%s,%t,%t,%t,%s,%s\n", repository.Id, repository.Name, projectID, repository.CreatedDate, repository.Size, repository.DefaultBranch, repository.URL, repository.RemoteUrl, repository.SSHUrl, repository.ValidRemoteUrls, repository.WebUrl, repository.IsDisabled, repository.IsFork, repository.IsInMaintenance, repository.ParentRepository.Id, repository.Project.Id)
		}
	}
	writeToFile(fileName, output)

}

func getBoards(organizationUrl string, authentication string, projectIDs []string, wg *sync.WaitGroup) {
	defer wg.Done()

	boardList := APIResults[boards]{
		Value: []boards{},
	}

	fileName := "boards.csv"
	output := "ID,Name,Project ID,URL\n"

	for _, projectID := range projectIDs {

		endpoint := EndPoint{
			resource:   "work/boards",
			parameters: "",
			urlBase:    organizationUrl + "/" + projectID,
			isGraph:    false,
		}

		boardList, err := getEndpointStruct(endpoint, boardList, authentication)
		if err != nil {
			log.Printf("getEndpointStruct(): %s - %s\n", endpoint.resource, err)
		}

		for _, board := range boardList.Value {
			output += fmt.Sprintf("%s,%s,%s,%s\n", board.Id, board.Name, projectID, board.URL)
		}

	}

	writeToFile(fileName, output)

}
