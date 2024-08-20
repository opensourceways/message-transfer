package dto

import (
	"github.com/opensourceways/go-gitee/gitee"
)

type GiteeIssueRaw struct {
	gitee.IssueEvent
	SigGroupName string `json:"sig_group_name"`
}

type GiteePushRaw struct {
	gitee.PushEvent
	SigGroupName string `json:"sig_group_name"`
}

type GiteePrRaw struct {
	gitee.PullRequestEvent
	SigGroupName string `json:"sig_group_name"`
}

type GiteeNoteRaw struct {
	gitee.NoteEvent
	SigGroupName string `json:"sig_group_name"`
}
