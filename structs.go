package main

type APIResults[T any] struct {
	Count int `json:"count"`
	Value []T `json:"value"`
}

type users struct {
	Descriptor    string `json:"descriptor"`
	DisplayName   string `json:"display_name"`
	PrincipalName string `json:"principal_name"`
	MailAddress   string `json:"mail_address"`
	SubjectKind   string `json:"subject_kind"`
	Domain        string `json:"domain"`
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

type teams struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	ProjectId   string `json:"projectid"`
	ProjectName string `json:"projectname"`
	URL         string `json:"url"`
	IdentityUrl string `json:"identityurl"`
	Identity    struct {
		CustomDisplayName   string   `json:"customdisplayname"`
		Descriptor          string   `json:"descriptor"`
		Id                  string   `json:"id"`
		IsActive            bool     `json:"isactive"`
		IsContainer         bool     `json:"iscontainer"`
		MasterId            string   `json:"masterid"`
		MemberIds           []string `json:"memberids"`
		MemberOf            []string `json:"memberof"`
		Members             []string `json:"members"`
		ProviderDisplayName string   `json:"providerDisplayName"`
		SubjectDescriptor   string   `json:"subjectdescriptor"`
		UniqueUserId        string   `json:"uniqueuserid"` //nolint:govet
	} `json:"identity"`
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
