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

func (raw *PrRaw) GetFollowUsers(events dto.CloudEvents) {
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
	events.SetExtension("followusers", strings.Join(utils.Difference(followUsers,
		raw.getTodoUsers()), ","))
}

func (raw *PrRaw) getTodoUsers() []string {
	var assignees []string
	for _, assignee := range raw.PullRequest.Assignees {
		assignees = append(assignees, assignee.UserName)
	}
	return assignees
}

func (raw *PrRaw) GetTodoUsers(events dto.CloudEvents) {
	events.SetExtension("todousers", strings.Join(raw.getTodoUsers(), ","))
	events.SetExtension("businessid", strconv.Itoa(int(raw.PullRequest.Id)))
}

func (raw *PrRaw) IsDone(events dto.CloudEvents) {
	events.SetExtension("isdone", *raw.State == "closed" || *raw.State == "merged")
}
