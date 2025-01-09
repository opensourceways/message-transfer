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

func (raw *IssueRaw) followUsers() []string {
	sigGroup, err := utils.GetRepoSigInfo(raw.Repository.Name)
	if err != nil {
		return []string{}
	}
	sigMaintainers, _, err := utils.GetMembersBySig(sigGroup)
	if err != nil {
		return []string{}
	}

	repo := strings.Split(raw.Repository.FullName, "/")
	repoAdmins, err := utils.GetAllAdmins(repo[0], repo[1])
	if err != nil {
		return []string{}
	}
	followUsers := slices.Concat(sigMaintainers, repoAdmins)
	followUsers = utils.Difference(followUsers, raw.todoUsers())
	followUsers = utils.Difference(followUsers, raw.applyUsers())
	return followUsers
}

func (raw *IssueRaw) GetFollowUsers(events dto.CloudEvents) {
	events.SetExtension("followusers", strings.Join(raw.followUsers(), ","))
}

func (raw *IssueRaw) todoUsers() []string {
	var todoUsers []string
	if raw.Issue.Assignee != nil {
		todoUsers = []string{(*raw.Issue.Assignee).UserName}
	}

	if raw.Assignee != nil {
		todoUsers = append(todoUsers, (*raw.Assignee).UserName)
	}
	todoUsers = utils.Difference(todoUsers, raw.applyUsers())
	return todoUsers
}

func (raw *IssueRaw) GetTodoUsers(events dto.CloudEvents) {
	events.SetExtension("todousers", strings.Join(raw.todoUsers(), ","))
	events.SetExtension("businessid", strconv.Itoa(int(raw.Issue.Id)))
}

func (raw *IssueRaw) applyUsers() []string {
	var applyUsers []string
	if raw.Sender != nil {
		applyUsers = append(applyUsers, raw.Sender.UserName)
	}
	return applyUsers
}

func (raw *IssueRaw) GetApplyUsers(events dto.CloudEvents) {
	events.SetExtension("applyusers", strings.Join(raw.applyUsers(), ","))
}

func (raw *IssueRaw) IsDone(events dto.CloudEvents) {
	events.SetExtension("isdone", slices.Contains([]string{"rejected", "closed"},
		raw.Issue.State))
}
