package main

type APIResults[T any] struct {
	Count int `json:"count"`
	Value []T `json:"value"`
}

type EndPoint struct {
	resource   string
	parameters string
	fileName   string
	headerRow  string
	urlBase    string
	isGraph    bool
}

type users struct {
	Descriptor    string `json:"descriptor"`
	Description   string `json:"description"`
	DisplayName   string `json:"displayname"`
	PrincipalName string `json:"principalname"`
	MailAddress   string `json:"mailaddress"`
	SubjectKind   string `json:"subjectkind"`
	Domain        string `json:"domain"`
	Origin        string `json:"origin"`
	OriginId      string `json:"originid"`
	URL           string `json:"url"`
}

type groups struct {
	Descriptor    string `json:"descriptor"`
	Description   string `json:"description"`
	DisplayName   string `json:"displayname"`
	PrincipalName string `json:"principalname"`
	MailAddress   string `json:"mailaddress"`
	SubjectKind   string `json:"subjectkind"`
	Domain        string `json:"domain"`
	Origin        string `json:"origin"`
	OriginId      string `json:"originid"`
	URL           string `json:"url"`
}

type project struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	State       string `json:"state"`
	Visibility  string `json:"visibility"`
	LastUpdate  string `json:"last_update"`
	URL         string `json:"url"`
}

type pipelines struct {
	Id            int    `json:"id"`
	Name          string `json:"name"`
	Folder        string `json:"folder"`
	Revision      int    `json:"revision"`
	URL           string `json:"url"`
	Configuration struct {
		Type string `json:"type"`
	}
}

type teams struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	ProjectId   string `json:"projectid"`
	ProjectName string `json:"projectname"`
	URL         string `json:"url"`
	IdentityUrl string `json:"identityurl"`
}

type repository struct {
	Id               string   `json:"id"`
	Name             string   `json:"name"`
	CreatedDate      string   `json:"createdDate"`
	Size             int      `json:"size"`
	DefaultBranch    string   `json:"defaultBranch"`
	URL              string   `json:"url"`
	RemoteUrl        string   `json:"remoteUrl"`
	SSHUrl           string   `json:"sshUrl"`
	ValidRemoteUrls  []string `json:"validRemoteUrls"`
	WebUrl           string   `json:"webUrl"`
	IsDisabled       bool     `json:"isDisabled"`
	IsFork           bool     `json:"isFork"`
	IsInMaintenance  bool     `json:"isInMaintenance"`
	ParentRepository struct {
		Id string `json:"id"`
	} `json:"parentRepository"`
	Project struct {
		Id string `json:"id"`
	} `json:"project"`
}

type boards struct {
	Id   string `json:"id"`
	Name string `json:"name"`
	URL  string `json:"url"`
}

type testPlan struct {
	Id              int    `json:"id"`
	Name            string `json:"name"`
	AreaPath        string `json:"areapath"`
	BuildDefinition struct {
		Id   int    `json:"id"`
		Name string `json:"name"`
	} `json:"buildDefinition"`
	Buildid     int    `json:"buildid"`
	Description string `json:"description"`
	Owner       struct {
		Id         int    `json:"id"`
		Descriptor string `json:"descriptor"`
	} `json:"owner"`
	PreviousBuildId              int `json:"previousBuildId"`
	ReleaseEnvironmentDefinition struct {
		DefinitionID            int    `json:"definitionid"`
		EnvironmentDefinitionId string `json:"environmentDefinitionId"`
	} `json:"releaseEnvironmentDefinition"`
	Revision  int `json:"revision"`
	RootSuite struct {
		Id   int    `json:"id"`
		Name string `json:"name"`
	} `json:"rootSuite"`
	StartDate string `json:"startDate"`
	EndDate   string `json:"endDate"`
	State     string `json:"state"`
	UpdatedBy struct {
		Id         int    `json:"id"`
		Descriptor string `json:"descriptor"`
	} `json:"updatedBy"`
	UpdatedDate          string `json:"updatedDate"`
	YamlReleaseReference struct {
		DefinitionID int    `json:"definitionid"`
		StagesToSkip string `json:"stagesToSkip"`
	} `json:"yamlReleaseReference"`
}
