/*
Copyright (c) Huawei Technologies Co., Ltd. 2024. All rights reserved
*/

// Package dto models dto of meeting
package dto

import (
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
	TodoId      string `json:"todoId"`
	TodoStatus  string `json:"todoStatus"`
}

func (raw CertificationRaw) GetRelateUsers(events CloudEvents) {
	events.SetExtension("relatedusers", "")
}

func (raw CertificationRaw) GetTodoUsers(events CloudEvents) {
	userId := raw.User
	userName, err := utils.GetUserNameById(userId)
	if err != nil {
		events.SetExtension("todousers", "")
		return
	}
	events.SetExtension("todousers", userName)
}

func (raw CertificationRaw) GetFollowUsers(events CloudEvents) {
	userIds := strings.Split(raw.User, ",")
	var userNames []string
	for _, userId := range userIds {
		userName, err := utils.GetUserNameById(userId)
		if err != nil {
			events.SetExtension("followusers", "")
			return
		}
		userNames = append(userNames, userName)
	}
	events.SetExtension("followusers", strings.Join(userNames, ","))
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
