/*
Copyright (c) Huawei Technologies Co., Ltd. 2024. All rights reserved
*/

package utils

import "time"

// Config represents the configuration for utils.
type Config struct {
	ConsumeSleepTime     int    `json:"consume_sleep_time"`
	GiteeAccessToken     string `json:"gitee_access_token"`
	GiteeCollaboratorUrl string `json:"gitee_collaborator_url"`
	GiteeWatcherUrl      string `json:"gitee_watcher_url"`
	GiteeContributorUrl  string `json:"gitee_contributor_url"`
	EulerRepoSigUrl      string `json:"euler_repo_sig_url"`
	EulerUserSigUrl      string `json:"euler_user_sig_url"`
	QuerySigInfo         string `json:"query_sig_info"`
	AuthorHost           string `json:"author_host"`
	Community            string `json:"community"`
	AppId                string `json:"app_id"`
	AppSecret            string `json:"app_secret"`
}

var config Config

// Init init utils config.
func Init(cfg *Config) error {
	config = *cfg
	return nil
}

// GetConsumeSleepTime get kafka consume sleep time.
func GetConsumeSleepTime() time.Duration {
	return time.Duration(config.ConsumeSleepTime) * time.Second
}
