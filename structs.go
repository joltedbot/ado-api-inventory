package main

type UserResults struct {
	Count int     `json:"count"`
	Value []users `json:"value"`
}

type users struct {
	Descriptor    string `json:"descriptor"`
	DisplayName   string `json:"display_name"`
	PrincipalName string `json:"principal_name"`
	MailAddress   string `json:"mail_address"`
	SubjectKind   string `json:"subject_kind"`
	Domain        string `json:"domain"`
}

type ProjectResults struct {
	Count int       `json:"count"`
	Value []project `json:"value"`
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

type TeamsResults struct {
	Count int     `json:"count"`
	Value []teams `json:"value"`
}

type teams struct {
	Id          string   `json:"id"`
	Name        string   `json:"name"`
	Description string   `json:"description"`
	ProjectId   string   `json:"projectid"`
	ProjectName string   `json:"projectname"`
	URL         string   `json:"url"`
	Identity    identity `json:"identity"`
	IdentityUrl string   `json:"identityurl"`
}

type identity struct {
	CustomDisplayName   string   `json:"customdisplayname"`
	Descriptor          string   `json:"descriptor"`
	Id                  string   `json:"id"`
	IsActive            bool     `json:"isactive"`
	IsContainer         bool     `json:"iscontainer"`
	MasterId            string   `json:"masterid"`
	MemberIds           []string `json:"memberids"`
	MemberOf            []string `json:"memberof"`
	Members             []string `json:"members"`
	ProviderDisplayName string   `json:"providerdisplayname"`
	SubjectDescriptor   string   `json:"subjectdescriptor"`
	UniqueUserId        string   `json:"uniqueuserid"` //nolint:govet
}
