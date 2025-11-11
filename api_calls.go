package main

import (
	"encoding/json"
	"fmt"
	"sync"
)

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
		println(err)
	}

	fileName := "graph-users.csv"
	output := "User,Email,Subject Kind,Principal Name,Domain\n"

	for _, user := range userList.Value {
		output += fmt.Sprintf("%s,%s,%s,%s,%s\n", user.DisplayName, user.MailAddress, user.SubjectKind, user.PrincipalName, user.Domain)
	}

	writeToFile(fileName, output)

}

func getProjects(organizationUrl string, authentication string, wg *sync.WaitGroup) {
	defer wg.Done()

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
		println(err)
	}

	fileName := "projects.csv"
	output := "Id,Name,Description,State,Visibility,LastUpdate,URL\n"

	for _, project := range projectList.Value {
		output += fmt.Sprintf("%s,%s,%s,%s,%s,%s,%s\n", project.Id, project.Name, project.Description, project.State, project.Visibility, project.LastUpdate, project.URL)
	}

	writeToFile(fileName, output)

}

func getRepositories(organizationUrl string, authentication string, wg *sync.WaitGroup) {
	defer wg.Done()

	projectList := APIResults[repository]{
		Value: []repository{},
	}

	endpoint := EndPoint{
		resource:   "git/repositories",
		parameters: "",
		urlBase:    organizationUrl,
		isGraph:    false,
	}

	projectList, err := getEndpointStruct(endpoint, projectList, authentication)
	if err != nil {
		println(err)
	}

	fileName := "repositories.csv"
	output := "Id,Name,Created Date,Size,Default Branch,URL,Remote URL,SSH URL,Valid Remote URLs,Web URL,Is Disabled,Is Fork,Is In Maintenance,Parent Repository ID,Project ID\n"

	for _, project := range projectList.Value {
		output += fmt.Sprintf("%s,%s,%s,%d,%s,%s,%s,%s,%s,%s,%t,%t,%t,%s,%s\n", project.Id, project.Name, project.CreatedDate, project.Size, project.DefaultBranch, project.URL, project.RemoteUrl, project.SSHUrl, project.ValidRemoteUrls, project.WebUrl, project.IsDisabled, project.IsFork, project.IsInMaintenance, project.ParentRepository.Id, project.Project.Id)
	}

	writeToFile(fileName, output)

}

func getTeams(organizationUrl string, authentication string, wg *sync.WaitGroup) {
	defer wg.Done()

	endpoint := "teams"
	parameters := "$expandIdentity=true"

	response, _, err := apiCall(endpoint, apiURL(false, organizationUrl, endpoint, parameters), "", authentication)
	if err != nil {
		println(err)
	}

	teamsList := APIResults[teams]{
		Value: []teams{},
	}

	err = json.Unmarshal([]byte(response), &teamsList)
	if err != nil {
		println(err)
	}

	teamsFileName := "teams.csv"
	teamsIdentityFileName := "teams-identities.csv"
	teamsOutput := "Id,Name,Description,Project ID,Project Name,URL,Identity Id,Identity URL\n"
	teamsIdentitiesOutput := "Custom Display Name,Descriptor,Team Id,Is Active,Is Container,Master Id,Member Ids,Member Of,Members,Provider Display Name,Subject Descriptor,Unique User Id\n"

	for _, team := range teamsList.Value {
		teamsOutput += fmt.Sprintf("%s,%s,%s,%s,%s,%s,%s,%s\n", team.Id, team.Name, team.Description, team.ProjectId, team.ProjectName, team.URL, team.Identity.Id, team.IdentityUrl)
		teamsIdentitiesOutput += fmt.Sprintf("%s,%s,%s,%t,%t,%s,%s,%s,%s,%s,%s,%s\n", team.Identity.CustomDisplayName, team.Identity.Descriptor, team.Id, team.Identity.IsActive, team.Identity.IsContainer, team.Identity.MasterId, team.Identity.MemberIds, team.Identity.MemberOf, team.Identity.Members, team.Identity.ProviderDisplayName, team.Identity.SubjectDescriptor, team.Identity.UniqueUserId)
	}

	writeToFile(teamsFileName, teamsOutput)
	writeToFile(teamsIdentityFileName, teamsIdentitiesOutput)

}
