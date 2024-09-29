/*
Copyright (c) Huawei Technologies Co., Ltd. 2024. All rights reserved
*/

// Package dto models dto of gitee
package dto

import (
	"github.com/opensourceways/go-gitee/gitee"
)

// GiteeIssueRaw gitee issue raw.
type GiteeIssueRaw struct {
	gitee.IssueEvent
	SigGroupName   string   `json:"sig_group_name"`
	SigMaintainers []string `json:"sig_maintainers"`
	RepoAdmins     []string `json:"repo_admins"`
}

// GiteePushRaw gitee push raw.
type GiteePushRaw struct {
	gitee.PushEvent
	SigGroupName   string   `json:"sig_group_name"`
	SigMaintainers []string `json:"sig_maintainers"`
	RepoAdmins     []string `json:"repo_admins"`
}

// GiteePrRaw gitee pr raw.
type GiteePrRaw struct {
	gitee.PullRequestEvent
	SigGroupName   string   `json:"sig_group_name"`
	SigMaintainers []string `json:"sig_maintainers"`
	RepoAdmins     []string `json:"repo_admins"`
}

// GiteeNoteRaw gitee note raw.
type GiteeNoteRaw struct {
	gitee.NoteEvent
	SigGroupName   string   `json:"sig_group_name"`
	SigMaintainers []string `json:"sig_maintainers"`
	RepoAdmins     []string `json:"repo_admins"`
}
