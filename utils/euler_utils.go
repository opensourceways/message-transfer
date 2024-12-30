/*
Copyright (c) Huawei Technologies Co., Ltd. 2024. All rights reserved
*/

package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
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

type ManagerTokenRequest struct {
	GrantType string `json:"grant_type"`
	AppId     string `json:"app_id"`
	AppSecret string `json:"app_secret"`
}

type ManagerTokenResponse struct {
	ManagerToken string `json:"token"`
}

type GetUserInfoResponse struct {
	Msg      string `json:"msg"`
	Code     int    `json:"code"`
	UserData `json:"data"`
}

type UserData struct {
	UserName string `json:"username"`
	Phone    string `json:"phone"`
	NickName string `json:"nickname"`
}

func JsonMarshal(t interface{}) ([]byte, error) {
	buffer := &bytes.Buffer{}
	enc := json.NewEncoder(buffer)
	enc.SetEscapeHTML(false)

	if err := enc.Encode(t); err != nil {
		return nil, err
	}
	return buffer.Bytes(), nil
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

// GetUserNameById get user name by id.
func GetUserNameById(userId string) (string, error) {
	url := fmt.Sprintf("%s/oneid/manager/getuserinfo?userId=%s", config.AuthorHost, userId)
	managerToken, err := getManagerToken(config.AppId, config.AppSecret)
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return "", err
	}
	req.Header.Add("token", managerToken)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	var data GetUserInfoResponse
	if err = json.Unmarshal(body, &data); err != nil {
		return "", err
	}
	if data.UserName == "" {
		return "", xerrors.Errorf("the user name is null")
	}
	return data.UserName, nil
}

func getManagerToken(appId string, appSecret string) (string, error) {
	url := fmt.Sprintf("%s/oneid/manager/token", config.AuthorHost)
	reqBody := ManagerTokenRequest{
		GrantType: "token",
		AppId:     appId,
		AppSecret: appSecret,
	}
	v, err := JsonMarshal(reqBody)
	if err != nil {
		return "", err
	}

	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(v))
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			return
		}
	}(resp.Body)
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	var data ManagerTokenResponse
	if err = json.Unmarshal(body, &data); err != nil {
		return "", err
	}
	return data.ManagerToken, nil
}
