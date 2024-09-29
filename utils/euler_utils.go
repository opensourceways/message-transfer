/*
Copyright (c) Huawei Technologies Co., Ltd. 2024. All rights reserved
*/

package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"golang.org/x/xerrors"
)

// RepoSig definition of repoSig.
type RepoSig struct {
	Sig string `json:"data"`
}

// SigInfo sig info.
type SigInfo struct {
	Data []SigData `json:"data"`
}

// SigData sig data.
type SigData struct {
	Maintainers []string `json:"maintainers"`
	Committers  []string `json:"committers"`
}

// GetRepoSigInfo get repo sig info.
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

// GetUserSigInfo get user sig info.
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

// GetMembersBySig get members by sig.
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

	if len(sigInfo.Data) == 0 {
		return []string{}, []string{}, xerrors.Errorf("no sig info data")
	}
	return sigInfo.Data[0].Maintainers, sigInfo.Data[0].Committers, nil
}
