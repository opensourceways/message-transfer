package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Collaborator struct {
	Id          int        `json:"id"`
	Login       string     `json:"login"`
	Name        string     `json:"name"`
	Type        string     `json:"type"`
	Permissions Permission `json:"permissions"`
}
type Contributor struct {
	Email         string `json:"email"`
	Name          string `json:"name"`
	Contributions int    `json:"contributions"`
}

type User struct {
	Id    int    `json:"id"`
	Login string `json:"login"`
	Name  string `json:"name"`
	Type  string `json:"type"`
}

type Permission struct {
	Pull  bool `json:"pull"`
	Push  bool `json:"push"`
	Admin bool `json:"admin"`
}

func (p *Permission) IsAdmin() bool {
	return p.Admin
}

func GetAllAdmins(owner, repo string) ([]string, error) {
	allCollaborators, err := fetchCollaborators(owner, repo)
	if err != nil {
		return nil, err
	}

	return filterAdmins(allCollaborators), nil
}

func fetchCollaborators(owner, repo string) ([]Collaborator, error) {
	var allCollaborators []Collaborator
	page := 1
	perPage := 100

	for {
		members, err := fetchCollaboratorsPage(owner, repo, page, perPage)
		if err != nil {
			return nil, err
		}

		allCollaborators = append(allCollaborators, members...)

		if len(members) < perPage {
			break
		}
		page++
	}
	return allCollaborators, nil
}

func fetchCollaboratorsPage(owner, repo string, page, perPage int) ([]Collaborator, error) {
	url := fmt.Sprintf(config.GiteeCollaboratorUrl, owner, repo, config.GiteeAccessToken, page, perPage)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var members []Collaborator
	if err = json.Unmarshal(body, &members); err != nil {
		return nil, err
	}

	return members, nil
}

func filterAdmins(collaborators []Collaborator) []string {
	var admins []string
	for _, collaborator := range collaborators {
		if collaborator.Permissions.IsAdmin() {
			admins = append(admins, collaborator.Login, collaborator.Name)
		}
	}
	return admins
}
