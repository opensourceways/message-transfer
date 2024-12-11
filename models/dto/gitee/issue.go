package gitee

import (
	"slices"
	"strconv"
	"strings"

	"github.com/opensourceways/go-gitee/gitee"
	"github.com/opensourceways/message-transfer/models/dto"
	"github.com/opensourceways/message-transfer/utils"
)

type IssueRaw struct {
	gitee.IssueEvent
}

func (raw *IssueRaw) ToCloudEventsByConfig(topic string) dto.CloudEvents {
	rawMap := dto.StructToMap(raw)
	return rawMap.ToCloudEventByConfig(topic)
}

func (raw *IssueRaw) GetRelateUsers(events dto.CloudEvents) {
	events.SetExtension("relatedusers", "")
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
	events.SetExtension("followusers", strings.Join(followUsers, ","))
}

func (raw *IssueRaw) GetTodoUsers(events dto.CloudEvents) {
	var todoUsers []string
	if raw.Issue.Assignee != nil {
		todoUsers = []string{(*raw.Issue.Assignee).UserName}
	}
	events.SetExtension("todousers", strings.Join(todoUsers, ","))
	events.SetExtension("businessid", strconv.Itoa(int(raw.Issue.Id)))
}

func (raw *IssueRaw) IsDone(events dto.CloudEvents) {
	events.SetExtension("isdone", slices.Contains([]string{"rejected", "closed"},
		raw.Issue.State))
}
