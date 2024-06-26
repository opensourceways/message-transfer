package dto

import (
	"github.com/opensourceways/go-gitee/gitee"
)

type GiteeIssueRaw struct {
	gitee.IssueEvent
}

type GiteePushRaw struct {
	gitee.PushEvent
}

type GiteePrRaw struct {
	gitee.PullRequestEvent
}

type GiteeNoteRaw struct {
	gitee.NoteEvent
}
