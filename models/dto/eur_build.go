/*
Copyright (c) Huawei Technologies Co., Ltd. 2024. All rights reserved
*/

// Package dto models dto of eur build.
package dto

import "strings"

// EurBuildMessageRaw eur build message raw.
type EurBuildMessageRaw struct {
	Body struct {
		User    string      `json:"user"`
		Copr    string      `json:"copr"`
		Owner   string      `json:"owner"`
		Pkg     interface{} `json:"pkg"`
		Build   int         `json:"build"`
		Chroot  string      `json:"chroot"`
		Version interface{} `json:"version"`
		Status  int         `json:"status"`
		IP      string      `json:"ip"`
		Who     string      `json:"who"`
		Pid     int         `json:"pid"`
		What    string      `json:"what"`
	} `json:"body"`
	Headers struct {
		OpenEulerMessagingSchema string `json:"openEuler_messaging_schema"`
		SentAt                   string `json:"sent-at"`
	} `json:"headers"`
	ID          string `json:"id"`
	Topic       string `json:"topic"`
	SourceGroup string `json:"source_group"`
}

func (raw *EurBuildMessageRaw) ToCloudEventsByConfig() CloudEvents {
	rawMap := StructToMap(raw)
	return rawMap.ToCloudEventByConfig("eur_build_raw")
}

func (raw *EurBuildMessageRaw) GetRelateUsers(events CloudEvents) {
	events.SetExtension("relatedusers", "")
}

func (raw *EurBuildMessageRaw) GetFollowUsers(events CloudEvents) {
	followUsers := []string{raw.Body.User, raw.Body.Owner}
	events.SetExtension("followusers", strings.Join(followUsers, ","))
}

func (raw *EurBuildMessageRaw) GetTodoUsers(events CloudEvents) {
	events.SetExtension("todousers", "")
}

func (raw *EurBuildMessageRaw) IsDone(events CloudEvents) {
	events.SetExtension("isdone", false)
}
