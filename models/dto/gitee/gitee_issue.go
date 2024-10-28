package gitee

import (
	"github.com/opensourceways/go-gitee/gitee"
	"github.com/opensourceways/message-transfer/models/dto"
)

type GiteeIssueRaw struct {
	gitee.IssueEvent
	//SigGroupName   string   `json:"sig_group_name"`
	//SigMaintainers []string `json:"sig_maintainers"`
	//RepoAdmins     []string `json:"repo_admins"`
}

func (raw *GiteeIssueRaw) ToCloudEventsByConfig() dto.CloudEvents {
	rawMap := dto.StructToMap(raw)
	return rawMap.ToCloudEventByConfig("gitee_issue_raw")
}

func (raw *GiteeIssueRaw) GetRelateUsers(events dto.CloudEvents) {
	releatedUsers := []string{}
	//TODO implement me

	//releatedUsers,_=utils.GetRepoSigInfo(raw.Repository.Name)
	events.SetExtension("releatedusers", releatedUsers)
}

func (raw *GiteeIssueRaw) GetFollowUsers(events dto.CloudEvents) {
	followUsers := []string{}
	//todo
	events.SetExtension("followUsers", followUsers)
}

func (raw *GiteeIssueRaw) GetTodoUsers(events dto.CloudEvents) {
	todoUsers := []string{}
	//todo
	events.SetExtension("todoUsers", todoUsers)
}
