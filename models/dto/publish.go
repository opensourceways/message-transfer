/*
Copyright (c) Huawei Technologies Co., Ltd. 2024. All rights reserved
*/

// Package dto models dto of eur build.
package dto

// PublishRaw eur build message raw.
type PublishRaw struct {
	Version string `json:"Version"`
	Url     string `json:"Url"`
}

func (raw *PublishRaw) ToCloudEventsByConfig(topic string) CloudEvents {
	rawMap := StructToMap(raw)
	return rawMap.ToCloudEventByConfig(topic)
}

func (raw *PublishRaw) GetRelateUsers(events CloudEvents) {
	events.SetExtension("relatedusers", "")
}

func (raw *PublishRaw) GetFollowUsers(events CloudEvents) {
	events.SetExtension("followusers", "")
}

func (raw *PublishRaw) GetTodoUsers(events CloudEvents) {
	events.SetExtension("todousers", "")
}

func (raw *PublishRaw) IsDone(events CloudEvents) {
	events.SetExtension("isdone", false)
}
