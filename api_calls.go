package main

import (
	"fmt"
	"log"
	"sync"
)

func getProjects(organizationUrl string, authentication string) ([]string, error) {
	endpoint := EndPoint{
		resource:   "projects",
		parameters: "",
		fileName:   "projects.csv",
		headerRow:  "Id,Name,Description,State,Visibility,LastUpdate,URL",
		urlBase:    organizationUrl,
		isGraph:    false,
	}

	var projectIDs []string

	err := fetchAndExport(endpoint, authentication, 0,
		func(p project) string {
			projectIDs = append(projectIDs, p.Id)
			return fmt.Sprintf("%s,%s,%s,%s,%s,%s,%s\n",
				p.Id, p.Name, p.Description, p.State, p.Visibility, p.LastUpdate, p.URL)
		},
	)

	if err != nil {
		return nil, err

	}

	return projectIDs, nil
}

func getTeams(organizationUrl string, authentication string, wg *sync.WaitGroup) {
	defer wg.Done()

	endpoint := EndPoint{
		resource:   "teams",
		parameters: "",
		fileName:   "teams.csv",
		headerRow:  "Id,Name,Description,Project ID,Project Name,URL,Identity Id,Identity URL",
		urlBase:    organizationUrl,
		isGraph:    false,
	}

	err := fetchAndExport(endpoint, authentication, 0,
		func(team teams) string {
			return fmt.Sprintf("%s,%s,%s,%s,%s,%s,%s,%s\n", team.Id, team.Name, team.Description, team.ProjectId, team.ProjectName, team.URL, team.Id, team.IdentityUrl)
		},
	)

	if err != nil {
		log.Printf("Error retrieving Teams data. Any output may be invalid or incomplete. Continuing anyway.")
	}

}

func getUsers(organizationUrl string, authentication string, wg *sync.WaitGroup) {
	defer wg.Done()

	endpoint := EndPoint{
		resource:   "graph/users",
		parameters: "",
		fileName:   "graph-users.csv",
		headerRow:  "Display Name, Descriptor, Description, Principal Name, Mail Address, Subject Kind, Domain, Origin, Origin ID, URL",
		urlBase:    organizationUrl,
		isGraph:    true,
	}

	err := fetchAndExport(endpoint, authentication, 0,
		func(user users) string {
			return fmt.Sprintf("%s,%s,%s,%s,%s,%s,%s,%s,%s,%s\n", user.DisplayName, user.Descriptor, user.Description, user.PrincipalName, user.MailAddress, user.SubjectKind, user.Domain, user.Origin, user.OriginId, user.URL)
		},
	)

	if err != nil {
		log.Printf("Error retrieving Users data. Any output may be invalid or incomplete. Continuing anyway.")
	}

}

func getGroups(organizationUrl string, authentication string, wg *sync.WaitGroup) {
	defer wg.Done()

	endpoint := EndPoint{
		resource:   "graph/groups",
		parameters: "",
		fileName:   "graph-groups.csv",
		headerRow:  "Display Name, Descriptor, Description, Principal Name, Mail Address, Subject Kind, Domain, Origin, Origin ID, URL",
		urlBase:    organizationUrl,
		isGraph:    true,
	}

	err := fetchAndExport(endpoint, authentication, 0,
		func(group groups) string {
			return fmt.Sprintf("%s,%s,%s,%s,%s,%s,%s,%s,%s,%s\n", group.DisplayName, group.Descriptor, group.Description, group.PrincipalName, group.MailAddress, group.SubjectKind, group.Domain, group.Origin, group.OriginId, group.URL)
		},
	)

	if err != nil {
		log.Printf("Error retrieving Groups data. Any output may be invalid or incomplete. Continuing anyway.")
	}

}

func getPipelines(organizationUrl string, authentication string, projectIDs []string, wg *sync.WaitGroup) {
	defer wg.Done()

	for index, projectID := range projectIDs {

		endpoint := EndPoint{
			resource:   "pipelines",
			parameters: "",
			fileName:   "pipelines.csv",
			headerRow:  "ID, Name, Folder, Revision, Project ID, URL, Configuration Type",
			urlBase:    organizationUrl + "/" + projectID,
			isGraph:    false,
		}

		err := fetchAndExport(endpoint, authentication, index,
			func(pipeline pipelines) string {
				return fmt.Sprintf("%d,%s,%s,%d,%s,%s,%s\n", pipeline.Id, pipeline.Name, pipeline.Folder, pipeline.Revision, projectID, pipeline.URL, pipeline.Configuration.Type)
			},
		)

		if err != nil {
			log.Printf("Error retrieving Pipeline data. Any output may be invalid or incomplete. Continuing anyway.")
		}

	}

}

func getRepositories(organizationUrl string, authentication string, projectIDs []string, wg *sync.WaitGroup) {
	defer wg.Done()

	for index, projectID := range projectIDs {
		endpoint := EndPoint{
			resource:   "git/repositories",
			parameters: "",
			fileName:   "repositories.csv",
			headerRow:  "Id,Name,Project ID,Created Date,Size,Default Branch,URL,Remote URL,SSH URL,Valid Remote URLs,Web URL,Is Disabled,Is Fork,Is In Maintenance,Parent Repository ID,Project ID",
			urlBase:    organizationUrl + "/" + projectID,
			isGraph:    false,
		}

		err := fetchAndExport(endpoint, authentication, index,
			func(repository repository) string {
				return fmt.Sprintf("%s,%s,%s,%s,%d,%s,%s,%s,%s,%s,%s,%t,%t,%t,%s,%s\n", repository.Id, repository.Name, projectID, repository.CreatedDate, repository.Size, repository.DefaultBranch, repository.URL, repository.RemoteUrl, repository.SSHUrl, repository.ValidRemoteUrls, repository.WebUrl, repository.IsDisabled, repository.IsFork, repository.IsInMaintenance, repository.ParentRepository.Id, repository.Project.Id)
			},
		)

		if err != nil {
			log.Printf("Error retrieving Repoisitory data. Any output may be invalid or incomplete. Continuing anyway.")
		}
	}

}

func getBoards(organizationUrl string, authentication string, projectIDs []string, wg *sync.WaitGroup) {
	defer wg.Done()

	for index, projectID := range projectIDs {

		endpoint := EndPoint{
			resource:   "work/boards",
			parameters: "",
			fileName:   "boards.csv",
			headerRow:  "ID,Name,Project ID,URL",
			urlBase:    organizationUrl + "/" + projectID,
			isGraph:    false,
		}

		err := fetchAndExport(endpoint, authentication, index,
			func(board boards) string {
				return fmt.Sprintf("%s,%s,%s,%s\n", board.Id, board.Name, projectID, board.URL)
			},
		)

		if err != nil {
			log.Printf("Error retrieving Repoisitory data. Any output may be invalid or incomplete. Continuing anyway.")
		}

	}

}
