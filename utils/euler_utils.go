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

type SigInfo struct {
	Data []SigData `json:"data"`
}

type SigData struct {
	Maintainers []string `json:"maintainers"`
	Committers  []string `json:"committers"`
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

func GetMembersBySig(sig string) ([]string, []string, error) {
	url := fmt.Sprintf(config.QuerySigInfo, sig)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, nil, err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, nil, err
	}

	var sigInfo SigInfo
	err = json.Unmarshal(body, &sigInfo)
	if err != nil {
		return nil, nil, err
	}

	return sigInfo.Data[0].Maintainers, sigInfo.Data[0].Committers, nil
}
