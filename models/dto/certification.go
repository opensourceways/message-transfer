/*
Copyright (c) Huawei Technologies Co., Ltd. 2024. All rights reserved
*/

// Package dto models dto of meeting
package dto

import (
	"strconv"
	"strings"
	"time"

	"github.com/opensourceways/message-transfer/utils"
)

// CertificationRaw openEuler meeting raw.
type CertificationRaw struct {
	Content     string `json:"content"`
	Time        int64  `json:"createTime"`
	RedirectUrl string `json:"redirectUrl"`
	Type        string `json:"type"`
	User        string `json:"user"`
	TodoId      int    `json:"todoId"`
	TodoStatus  string `json:"todoStatus"`
}

func (raw CertificationRaw) GetRelateUsers(events CloudEvents) {
	events.SetExtension("relatedusers", "")
}

func (raw CertificationRaw) todoUsers() []string {
	var todoUsers []string
	if raw.Type == "todo" {
		return []string{}
	}
	userId := raw.User
	userName, err := utils.GetUserNameById(userId)
	if err != nil {
		return []string{}
	}
	todoUsers = []string{userName}
	return todoUsers
}

func (raw CertificationRaw) GetTodoUsers(events CloudEvents) {
	events.SetExtension("todousers", strings.Join(raw.todoUsers(), ","))
	events.SetExtension("businessid", strconv.Itoa(raw.TodoId))
}

func (raw CertificationRaw) followUsers() []string {
	var followUsers []string
	if raw.Type != "notice" {
		return []string{}
	}
	userIds := strings.Split(raw.User, ",")
	var userNames []string
	for _, userId := range userIds {
		userName, err := utils.GetUserNameById(userId)
		if err != nil {
			return []string{}
		}
		userNames = append(userNames, userName)
	}
	followUsers = userNames
	followUsers = utils.Difference(followUsers, raw.todoUsers())
	return followUsers
}

func (raw CertificationRaw) GetFollowUsers(events CloudEvents) {
	events.SetExtension("followusers", strings.Join(raw.followUsers(), ","))
}

func (raw CertificationRaw) applyUsers() []string {
	return []string{}
}

func (raw CertificationRaw) GetApplyUsers(events CloudEvents) {
	events.SetExtension("applyusers", "")
}

func (raw CertificationRaw) ToCloudEventsByConfig(topic string) CloudEvents {
	rawMap := StructToMap(raw)
	return rawMap.ToCloudEventByConfig(topic)
}

func (raw CertificationRaw) IsDone(events CloudEvents) {
	seconds := raw.Time / 1000
	nanoseconds := (raw.Time % 1000) * 1000000
	t := time.Unix(seconds, nanoseconds)
	events.SetTime(t)
	events.SetExtension("isdone", raw.TodoStatus == "done")
}
