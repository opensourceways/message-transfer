package gitee

import (
	"regexp"
	"slices"
	"strings"

	"github.com/opensourceways/go-gitee/gitee"
	"github.com/sirupsen/logrus"

	"github.com/opensourceways/message-transfer/models/dto"
	"github.com/opensourceways/message-transfer/utils"
)

// NoteRaw gitee note raw.
type NoteRaw struct {
	gitee.NoteEvent
}

func IsBot(raw *NoteRaw) bool {
	sendUser := raw.Sender.Login
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
	logrus.Infof(*raw.NoteableType, raw.Issue, raw.PullRequest)
	if *raw.NoteableType == "PullRequest" {
		if raw.PullRequest == nil || raw.PullRequest.User == nil {
			return ""
		}
		return raw.PullRequest.User.UserName
	} else {
		if raw.Issue == nil || raw.Issue.User == nil {
			return ""
		}
		return raw.Issue.User.UserName
	}
}

func (raw *NoteRaw) ToCloudEventsByConfig() dto.CloudEvents {
	rawMap := dto.StructToMap(raw)
	return rawMap.ToCloudEventByConfig("gitee_note_raw")
}

func (raw *NoteRaw) GetRelateUsers(events dto.CloudEvents) {
	if !IsBot(raw) {
		mention, owner := GetMentionedUsers(raw), GetOwner(raw)
		events.SetExtension("releatedUsers", append(mention, owner))
	}
}

func (raw *NoteRaw) GetFollowUsers(events dto.CloudEvents) {
	mention := GetMentionedUsers(raw)
	owner := GetOwner(raw)
	repo := strings.Split(raw.Repository.FullName, "/")
	repoAdmins, err := utils.GetAllAdmins(repo[0], repo[1])
	if err != nil {
		return
	}

	mentionAndOwner := append(mention, owner)
	if IsBot(raw) {
		events.SetExtension("followUsers", mentionAndOwner)
	} else {
		events.SetExtension("followUsers", utils.Difference(repoAdmins, mentionAndOwner))
	}
}

func (raw *NoteRaw) GetTodoUsers(events dto.CloudEvents) {
	events.SetExtension("todoUsers", []string{})
}

func (raw *NoteRaw) IsDone(events dto.CloudEvents) {
	return
}
