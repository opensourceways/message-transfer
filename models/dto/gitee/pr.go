package gitee

import (
	"slices"
	"strconv"
	"strings"

	"github.com/opensourceways/go-gitee/gitee"

	"github.com/opensourceways/message-transfer/models/dto"
	"github.com/opensourceways/message-transfer/utils"
)

// PrRaw gitee pr raw.
type PrRaw struct {
	gitee.PullRequestEvent
}

func (raw *PrRaw) ToCloudEventsByConfig(topic string) dto.CloudEvents {
	rawMap := dto.StructToMap(raw)
	return rawMap.ToCloudEventByConfig(topic)
}

func (raw *PrRaw) GetRelateUsers(events dto.CloudEvents) {
	events.SetExtension("relatedusers", "")
}

func (raw *PrRaw) followUsers() []string {
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

func (raw *PrRaw) GetFollowUsers(events dto.CloudEvents) {
	events.SetExtension("followusers", strings.Join(raw.followUsers(), ","))
}

func (raw *PrRaw) todoUsers() []string {
	var todoUsers []string
	for _, assignee := range raw.PullRequest.Assignees {
		todoUsers = append(todoUsers, assignee.UserName)
	}
	todoUsers = utils.Difference(todoUsers, raw.applyUsers())
	return todoUsers
}

func (raw *PrRaw) GetTodoUsers(events dto.CloudEvents) {
	events.SetExtension("todousers", strings.Join(raw.todoUsers(), ","))
	events.SetExtension("businessid", strconv.Itoa(int(raw.PullRequest.Id)))
}

func (raw *PrRaw) applyUsers() []string {
	var applyUsers []string
	if raw.Sender != nil {
		applyUsers = []string{raw.Sender.UserName}
	}
	return applyUsers
}

func (raw *PrRaw) GetApplyUsers(events dto.CloudEvents) {
	events.SetExtension("applyusers", strings.Join(raw.applyUsers(), ","))
}

func (raw *PrRaw) IsDone(events dto.CloudEvents) {
	events.SetExtension("isdone", *raw.State == "closed" || *raw.State == "merged")
}
