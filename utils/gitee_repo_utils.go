package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
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
	var allCollaborators []Collaborator
	page := 1
	perPage := 100

	var totalCount int

	for {
		url := fmt.Sprintf(config.GiteeCollaboratorUrl, owner, repo, config.GiteeAccessToken, page, perPage)
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return nil, err
		}

		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			return nil, err
		}

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}

		var members []Collaborator
		err = json.Unmarshal(body, &members)
		if err != nil {
			return nil, err
		}

		allCollaborators = append(allCollaborators, members...)

		if totalCount == 0 {
			totalCount, err = strconv.Atoi(resp.Header.Get("total_count"))
			if err != nil {
				return nil, err
			}
		}
		if len(members) < perPage {
			break
		}
		page++
		err = resp.Body.Close()
		if err != nil {
			return nil, err
		}
	}

	var admins []string
	for _, collaborator := range allCollaborators {
		if collaborator.Permissions.IsAdmin() {
			admins = append(admins, collaborator.Login, collaborator.Name)
		}
	}
	return admins, nil
}
