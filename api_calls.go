package main

import (
	"fmt"
	"log"
	"sync"
)

func getProjects(organizationUrl string, authentication string) ([]string, error) {
	endpoint := EndPoint{
		urlBase:      "https://dev.azure.com",
		resource:     "projects",
		parameters:   "",
		fileName:     "projects.csv",
		headerRow:    "Id,Name,Description,State,Visibility,LastUpdate,URL",
		organization: organizationUrl,
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
		urlBase:      "https://dev.azure.com",
		resource:     "teams",
		parameters:   "",
		fileName:     "teams.csv",
		headerRow:    "Id,Name,Description,Project ID,Project Name,URL,Identity Id,Identity URL",
		organization: organizationUrl,
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
		urlBase:      "https://vssps.dev.azure.com",
		resource:     "graph/users",
		parameters:   "",
		fileName:     "graph-users.csv",
		headerRow:    "Display Name, Descriptor, Description, Principal Name, Mail Address, Subject Kind, Domain, Origin, Origin ID, URL",
		organization: organizationUrl,
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
		urlBase:      "https://vssps.dev.azure.com",
		resource:     "graph/groups",
		parameters:   "",
		fileName:     "graph-groups.csv",
		headerRow:    "Display Name, Descriptor, Description, Principal Name, Mail Address, Subject Kind, Domain, Origin, Origin ID, URL",
		organization: organizationUrl,
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
			urlBase:      "https://dev.azure.com",
			resource:     "pipelines",
			parameters:   "",
			fileName:     "pipelines.csv",
			headerRow:    "ID, Name, Folder, Revision, Project ID, URL, Configuration Type",
			organization: organizationUrl + "/" + projectID,
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
			urlBase:      "https://dev.azure.com",
			resource:     "git/repositories",
			parameters:   "",
			fileName:     "repositories.csv",
			headerRow:    "Id,Name,Project ID,Created Date,Size,Default Branch,URL,Remote URL,SSH URL,Valid Remote URLs,Web URL,Is Disabled,Is Fork,Is In Maintenance,Parent Repository ID,Project ID",
			organization: organizationUrl + "/" + projectID,
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
			urlBase:      "https://dev.azure.com",
			resource:     "work/boards",
			parameters:   "",
			fileName:     "boards.csv",
			headerRow:    "ID,Name,Project ID,URL",
			organization: organizationUrl + "/" + projectID,
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

func getTestPlans(organizationUrl string, authentication string, projectIDs []string, wg *sync.WaitGroup) {
	defer wg.Done()

	for index, projectID := range projectIDs {

		endpoint := EndPoint{
			urlBase:      "https://dev.azure.com",
			resource:     "testplan/plans",
			parameters:   "filterActivePlans=true",
			fileName:     "testplans.csv",
			headerRow:    "Id,Name,Area Path,Build Definition ID,Build Definition Name,Build ID,Description,Owner.ID,Owner Descriptor,Previous Build ID,Release Environment Definition ID,Environment Definition ID,Revision,Root Suite ID,Root Suite Name,Start Date,End Date,State,Updated By ID, Updated By Descriptor,Updated Date,Yaml Release Reference Definition ID,Yaml Release Reference Definition Stages To Skip",
			organization: organizationUrl + "/" + projectID,
		}

		err := fetchAndExport(endpoint, authentication, index,
			func(testplan testPlan) string {
				return fmt.Sprintf("%d,%s,%s,%d,%s,%d,%s,%d,%s,%d,%d,%s,%d,%d,%s,%s,%s,%s,%d,%s,%s,%d,%s\n", testplan.Id, testplan.Name, testplan.AreaPath, testplan.BuildDefinition.Id, testplan.BuildDefinition.Name, testplan.Buildid, testplan.Description, testplan.Owner.Id, testplan.Owner.Descriptor, testplan.PreviousBuildId, testplan.ReleaseEnvironmentDefinition.DefinitionID, testplan.ReleaseEnvironmentDefinition.EnvironmentDefinitionId, testplan.Revision, testplan.RootSuite.Id, testplan.RootSuite.Name, testplan.StartDate, testplan.EndDate, testplan.State, testplan.UpdatedBy.Id, testplan.UpdatedBy.Descriptor, testplan.UpdatedDate, testplan.YamlReleaseReference.DefinitionID, testplan.YamlReleaseReference.StagesToSkip)
			},
		)

		if err != nil {
			log.Printf("Error retrieving Repoisitory data. Any output may be invalid or incomplete. Continuing anyway.")
		}

	}

}

func getWiki(organizationUrl string, authentication string, wg *sync.WaitGroup) {
	defer wg.Done()

	endpoint := EndPoint{
		urlBase:      "https://dev.azure.com/",
		resource:     "wiki/wikis",
		parameters:   "",
		fileName:     "wiki.csv",
		headerRow:    "ID,Name,Is Disabled,Mapped Path,Project ID,Remote URL,Repository ID,Type,URL",
		organization: organizationUrl,
	}

	err := fetchAndExport(endpoint, authentication, 0,
		func(wiki wikis) string {
			return fmt.Sprintf("%s,%s,%t,%s,%s,%s,%s,%s,%s\n", wiki.Id, wiki.Name, wiki.IsDisabled, wiki.MappedPath, wiki.Projectid, wiki.RemoteUrl, wiki.RepositoryID, wiki.Type, wiki.URL)
		},
	)

	if err != nil {
		log.Printf("Error retrieving Repoisitory data. Any output may be invalid or incomplete. Continuing anyway.")
	}
}

func getArtifactFeeds(organizationUrl string, authentication string, projectIDs []string, wg *sync.WaitGroup) {
	defer wg.Done()

	for index, projectID := range projectIDs {

		endpoint := EndPoint{
			urlBase:      "https://feeds.dev.azure.com",
			resource:     "packaging/feeds",
			parameters:   "",
			fileName:     "artifact-feeds.csv",
			headerRow:    "Id,Name,BadgesEnabled,Capabilities,DefaultViewId,DeletedDate,Description,FullyQualifiedId,FullyQualifiedName,HideDeletedPackageVersions,IsEnabled,IsReadOnly,PermanentDeletedDate,ScheduledPermanentDeleteDate,UpstreamEnabled,UpstreamEnabledChangedDate,Upstream Sources,View Id,View Name,View Type,View Visibility,View URL,ViewId,ViewName,URL",
			organization: organizationUrl + "/" + projectID,
		}

		err := fetchAndExport(endpoint, authentication, index,

			func(feed artifactFeeds) string {
				var upstream string
				for index, stream := range feed.UpstreamSources {
					if index == 0 {
						upstream = upstream + stream.Name
					} else {
						upstream = upstream + " | " + stream.Name
					}

				}
				return fmt.Sprintf("%s,%s,%t,%s,%s,%s,%s,%s,%s,%t,%t,%t,%s,%s,%t,%s,%s,%s,%s,%s,%s,%s,%s,%s,%s\n", feed.Id, feed.Name, feed.BadgesEnabled, feed.Capabilities, feed.DefaultViewId, feed.DeletedDate, feed.Description, feed.FullyQualifiedId, feed.FullyQualifiedName, feed.HideDeletedPackageVersions, feed.IsEnabled, feed.IsReadOnly, feed.PermanentDeletedDate, feed.ScheduledPermanentDeleteDate, feed.UpstreamEnabled, feed.UpstreamEnabledChangedDate, upstream, feed.View.Id, feed.View.Name, feed.View.Type, feed.View.Visibility, feed.View.URL, feed.ViewId, feed.ViewName, feed.URL)
			},
		)

		if err != nil {
			log.Printf("Error retrieving Repoisitory data. Any output may be invalid or incomplete. Continuing anyway.")
		}

	}

}
