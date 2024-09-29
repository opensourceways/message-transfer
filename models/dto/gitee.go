package dto

import (
	"github.com/opensourceways/go-gitee/gitee"
)

type GiteeIssueRaw struct {
	gitee.IssueEvent
	SigGroupName   string   `json:"sig_group_name"`
	SigMaintainers []string `json:"sig_maintainers"`
	RepoAdmins     []string `json:"repo_admins"`
}

type GiteePushRaw struct {
	gitee.PushEvent
	SigGroupName   string   `json:"sig_group_name"`
	SigMaintainers []string `json:"sig_maintainers"`
	RepoAdmins     []string `json:"repo_admins"`
}

type GiteePrRaw struct {
	gitee.PullRequestEvent
	SigGroupName   string   `json:"sig_group_name"`
	SigMaintainers []string `json:"sig_maintainers"`
	RepoAdmins     []string `json:"repo_admins"`
}

type GiteeNoteRaw struct {
	gitee.NoteEvent
	SigGroupName   string   `json:"sig_group_name"`
	SigMaintainers []string `json:"sig_maintainers"`
	RepoAdmins     []string `json:"repo_admins"`
}
