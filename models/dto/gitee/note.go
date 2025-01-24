package gitee

import (
	"regexp"
	"slices"
	"strings"

	"github.com/opensourceways/go-gitee/gitee"
	"github.com/opensourceways/message-transfer/models/dto"
	"github.com/opensourceways/message-transfer/utils"
)

// NoteRaw gitee note raw.
type NoteRaw struct {
	gitee.NoteEvent
}

func IsBot(raw *NoteRaw) bool {
	sendUser := (*raw.Sender).Login
	botNames := []string{"ci-robot", "openeuler-ci-bot", "openeuler-sync-bot"}
	return slices.Contains(botNames, sendUser)
}

func GetMentionedUsers(raw *NoteRaw) []string {
	noteContent := raw.Comment.Body
	regex := regexp.MustCompile(`@(\w+)`) // 匹配以@开头的用户名
	matches := regex.FindAllStringSubmatch(noteContent, -1)

	var usernames []string
	for _, match := range matches {
		if len(match) > 1 {
			usernames = append(usernames, match[1]) // 添加匹配的用户名
		}
	}
	return usernames
}

func GetOwner(raw *NoteRaw) string {
	if *raw.NoteableType == "PullRequest" {
		if raw.PullRequest == nil || raw.PullRequest.User == nil {
			return ""
		}
		return (*raw.PullRequest.User).UserName
	} else {
		if raw.Issue == nil || raw.Issue.User == nil {
			return ""
		}
		return (*raw.Issue.User).UserName
	}
}

func (raw *NoteRaw) ToCloudEventsByConfig(topic string) dto.CloudEvents {
	rawMap := dto.StructToMap(raw)
	return rawMap.ToCloudEventByConfig(topic)
}

func (raw *NoteRaw) relatedUsers() []string {
	var relatedUsers []string
	if !IsBot(raw) {
		mention, owner := GetMentionedUsers(raw), GetOwner(raw)
		relatedUsers = append(mention, owner)
	} else {
		relatedUsers = []string{}
	}
	return relatedUsers
}

func (raw *NoteRaw) GetRelateUsers(events dto.CloudEvents) {
	events.SetExtension("relatedusers", strings.Join(raw.relatedUsers(), ","))
}

func (raw *NoteRaw) followUsers() []string {
	var followUsers []string
	mention := GetMentionedUsers(raw)
	owner := GetOwner(raw)
	repo := strings.Split(raw.Repository.FullName, "/")
	repoAdmins, err := utils.GetAllAdmins(repo[0], repo[1])
	if err != nil {
		return []string{}
	}
	mentionAndOwner := append(mention, owner)

	if *raw.NoteableType == "PullRequest" {
		for _, assignee := range raw.PullRequest.Assignees {
			mentionAndOwner = append(mentionAndOwner, assignee.UserName)
		}
	} else if *raw.NoteableType == "Issue" {
		if raw.Issue.Assignee != nil {
			mentionAndOwner = append(mentionAndOwner, (*raw.Issue.Assignee).UserName)
		}
	}

	if IsBot(raw) {
		followUsers = mentionAndOwner
	} else {
		followUsers = utils.Difference(repoAdmins, raw.relatedUsers())
	}
	return followUsers
}

func (raw *NoteRaw) GetFollowUsers(events dto.CloudEvents) {
	events.SetExtension("followusers", strings.Join(raw.followUsers(), ","))
}

func (raw *NoteRaw) GetTodoUsers(events dto.CloudEvents) {
	events.SetExtension("todousers", "")
}

func (raw *NoteRaw) IsDone(events dto.CloudEvents) {
	events.SetExtension("isdone", false)
}
