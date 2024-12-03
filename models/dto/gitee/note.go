package gitee

import (
	"regexp"
	"slices"
	"strings"

	"github.com/opensourceways/go-gitee/gitee"
	"github.com/opensourceways/message-transfer/models/dto"
	"github.com/opensourceways/message-transfer/utils"
	"github.com/sirupsen/logrus"
)

// NoteRaw gitee note raw.
type NoteRaw struct {
	gitee.NoteEvent
}

func IsBot(raw *NoteRaw) bool {
	sendUser := (*raw.Sender).Login
	botNames := []string{"ci-robot", "openeuler-ci-bot", "openeuler-sync-bot"}
	logrus.Infof("isbot: %v", slices.Contains(botNames, sendUser))
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

func (raw *NoteRaw) ToCloudEventsByConfig() dto.CloudEvents {
	rawMap := dto.StructToMap(raw)
	return rawMap.ToCloudEventByConfig("gitee_note_raw")
}

func (raw *NoteRaw) GetRelateUsers(events dto.CloudEvents) {
	if !IsBot(raw) {
		mention, owner := GetMentionedUsers(raw), GetOwner(raw)
		events.SetExtension("relatedusers", append(mention, owner))
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
	logrus.Infof("isbot : %v", IsBot(raw))
	if IsBot(raw) {
		events.SetExtension("followusers", mentionAndOwner)
		logrus.Infof("mentionAndOwner is %v", mentionAndOwner)
		logrus.Infof("1111 followusers: %v", events.Extensions()["followusers"])
	} else {
		events.SetExtension("followusers", utils.Difference(repoAdmins, mentionAndOwner))
		logrus.Infof("2222 followusers: %v", events.Extensions()["followusers"])
	}
}

func (raw *NoteRaw) GetTodoUsers(events dto.CloudEvents) {
	events.SetExtension("todousers", []string{})
}

func (raw *NoteRaw) IsDone(events dto.CloudEvents) {
	return
}
