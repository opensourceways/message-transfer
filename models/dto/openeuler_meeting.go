/*
Copyright (c) Huawei Technologies Co., Ltd. 2024. All rights reserved
*/

// Package dto models dto of meeting
package dto

import (
	"fmt"
	"time"

	"github.com/opensourceways/message-transfer/utils"
	"github.com/sirupsen/logrus"
)

// OpenEulerMeetingRaw openEuler meeting raw.
type OpenEulerMeetingRaw struct {
	Action string `json:"action"`
	Msg    struct {
		Id        int    `json:"id"`
		Topic     string `json:"topic"`
		Community string `json:"community"`
		GroupName string `json:"group_name"`
		Sponsor   string `json:"sponsor"`
		Date      string `json:"date"`
		Start     string `json:"start"`
		End       string `json:"end"`
		Duration  string `json:"duration"`
		Agenda    string `json:"agenda"`
		Etherpad  string `json:"etherpad"`
		Emaillist string `json:"emaillist"`
		HostId    string `json:"host_id"`
		Mid       string `json:"mid"`
		Mmid      string `json:"mmid"`
		JoinUrl   string `json:"join_url"`
		IsDelete  int    `json:"is_delete"`
		StartUrl  string `json:"start_url"`
		Timezone  string `json:"timezone"`
		User      int    `json:"user"`
		Group     int    `json:"group"`
		Mplatform string `json:"mplatform"`
	} `json:"msg"`
}

func (raw OpenEulerMeetingRaw) GetRelateUsers(events CloudEvents) {
	events.SetExtension("releatedusers", []string{})
}

func (raw OpenEulerMeetingRaw) GetTodoUsers(events CloudEvents) {
	sigMaintainers, _, err := utils.GetMembersBySig(raw.Msg.GroupName)
	if err != nil {
		logrus.Errorf("get members by sig failed, err:%v", err)
		events.SetExtension("todousers", []string{})
		return
	}
	events.SetExtension("todousers", sigMaintainers)
	events.SetExtension("businessid", raw.Msg.Id)
}

func (raw OpenEulerMeetingRaw) GetFollowUsers(events CloudEvents) {
	events.SetExtension("followusers", []string{})
}

func (raw OpenEulerMeetingRaw) ToCloudEventsByConfig() CloudEvents {
	rawMap := StructToMap(raw)
	return rawMap.ToCloudEventByConfig("openEuler_meeting_raw")
}

func (raw OpenEulerMeetingRaw) IsDone(events CloudEvents) {

	if raw.Action == "delete_meeting" {
		events.SetExtension("isdone", true)
		return
	}
	layout := "2006-01-0215:04"

	meetingEndTime, err := time.Parse(layout, raw.Msg.Date+raw.Msg.End)
	if err != nil {
		fmt.Println("Error parsing end time:", err)
		return
	}

	// 获取当前时间
	currentTime := time.Now()

	// 判断当前时间是否过期
	if currentTime.After(meetingEndTime) {
		events.SetExtension("isdone", true)
	}
}
