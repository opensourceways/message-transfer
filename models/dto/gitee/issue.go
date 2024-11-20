package gitee

import (
	"slices"
	"strings"

	"github.com/opensourceways/go-gitee/gitee"
	"github.com/opensourceways/message-transfer/models/dto"
	"github.com/opensourceways/message-transfer/utils"
)

type IssueRaw struct {
	gitee.IssueEvent
}

func (raw *IssueRaw) ToCloudEventsByConfig() dto.CloudEvents {
	rawMap := dto.StructToMap(raw)
	return rawMap.ToCloudEventByConfig("gitee_issue_raw")
}

func (raw *IssueRaw) GetRelateUsers(events dto.CloudEvents) {
	events.SetExtension("releatedusers", []string{})
}

func (raw *IssueRaw) GetFollowUsers(events dto.CloudEvents) {
	sigGroup, err := utils.GetRepoSigInfo(raw.Repository.Name)
	if err != nil {
		return
	}
	sigMaintainers, _, err := utils.GetMembersBySig(sigGroup)
	if err != nil {
		return
	}

	repo := strings.Split(raw.Repository.FullName, "/")
	repoAdmins, err := utils.GetAllAdmins(repo[0], repo[1])
	if err != nil {
		return
	}
	followUsers := slices.Concat(sigMaintainers, repoAdmins)
	events.SetExtension("followusers", followUsers)
}

func (raw *IssueRaw) GetTodoUsers(events dto.CloudEvents) {
	todoUsers := []string{raw.Issue.Assignee.UserName}
	events.SetExtension("todoUsers", todoUsers)
	events.SetExtension("businessid", raw.Issue.Id)
}

func (raw *IssueRaw) IsDone(events dto.CloudEvents) {
	events.SetExtension("isdone", slices.Contains([]string{"rejected", "closed"},
		raw.Issue.State))
}
