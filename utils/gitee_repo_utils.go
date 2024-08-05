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
type Watcher struct {
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

type Permission struct {
	Pull  bool `json:"pull"`
	Push  bool `json:"push"`
	Admin bool `json:"admin"`
}

func (p *Permission) IsAdmin() bool {
	return p.Admin
}

const (
	accessToken      = "****"
	collaboratorsUrl = "https://gitee.com/api/v5/repos/%s/%s/collaborators?access_token=%s&page=%d&per_page=%d"
	watchersUrl      = "https://gitee.com/api/v5/repos/%s/%s/subscribers?access_token=%s&page=%d&per_page=%d"
	contributorsUrl  = "https://gitee.com/api/v5/repos/%s/%s/contributors?access_token=%s&type=authors"
)

func GetAllCollaborators(owner, repo string) ([]Collaborator, error) {
	var allCollaborators []Collaborator
	page := 1
	perPage := 100

	var totalCount int

	for {
		url := fmt.Sprintf(collaboratorsUrl, owner, repo, accessToken, page, perPage)
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
		err = json.Unmarshal(body, &members)
		if err != nil {
			return nil, err
		}

		allCollaborators = append(allCollaborators, members...)

		if totalCount == 0 {
			totalCount, _ = strconv.Atoi(resp.Header.Get("total_count"))
		}

		if len(members) < perPage {
			break
		}
		page++
	}

	return allCollaborators, nil
}

func GetAllWatchers(owner, repo string) ([]Watcher, error) {
	var allWatchers []Watcher
	page := 1
	perPage := 100

	var totalCount int

	for {
		url := fmt.Sprintf(watchersUrl, owner, repo, accessToken, page, perPage)
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

		var members []Watcher
		err = json.Unmarshal(body, &members)
		if err != nil {
			return nil, err
		}

		allWatchers = append(allWatchers, members...)

		if totalCount == 0 {
			totalCount, _ = strconv.Atoi(resp.Header.Get("total_count"))
		}

		if len(members) < perPage {
			break
		}
		page++
	}

	return allWatchers, nil
}

func GetAllContributors(owner, repo string) ([]Contributor, error) {
	var allContributors []Contributor
	url := fmt.Sprintf(contributorsUrl, owner, repo, accessToken)
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

	var members []Contributor
	err = json.Unmarshal(body, &members)
	if err != nil {
		return nil, err
	}

	allContributors = append(allContributors, members...)

	return allContributors, nil
}
