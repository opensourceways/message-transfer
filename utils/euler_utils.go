package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type RepoSig struct {
	Sig string `json:"data"`
}

func GetRepoSigInfo(repo string) (string, error) {
	url := fmt.Sprintf(config.EulerRepoSigUrl, repo)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var repoSig RepoSig
	err = json.Unmarshal(body, &repoSig)
	if err != nil {
		return "", err
	}

	return repoSig.Sig, nil
}

func GetUserSigInfo(userName string) (string, error) {
	url := fmt.Sprintf(config.EulerUserSigUrl, userName)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var repoSig RepoSig
	err = json.Unmarshal(body, &repoSig)
	if err != nil {
		return "", err
	}

	return repoSig.Sig, nil
}
