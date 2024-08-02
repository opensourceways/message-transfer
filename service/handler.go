package service

import (
	"encoding/json"
	"github.com/opensourceways/message-transfer/common/kafka"
	"github.com/opensourceways/message-transfer/config"
	"github.com/opensourceways/message-transfer/models/dto"
)

func handle(raw dto.Raw, cfg kafka.ConsumeConfig) error {
	event := raw.ToCloudEventByConfig(cfg.Topic)
	if event.ID() == "" {
		return nil
	}
	kafkaSendErr := kafka.SendMsg(cfg.Publish, &event)
	if kafkaSendErr != nil {
		return kafkaSendErr
	}
	event.SaveDb()
	return nil
}

func GiteeIssueHandle(payload []byte, _ map[string]string) error {
	var raw dto.GiteeIssueRaw
	msgBodyErr := json.Unmarshal(payload, &raw)
	if msgBodyErr != nil {
		return msgBodyErr
	}
	rawMap := dto.StructToMap(raw)
	return handle(rawMap, config.GiteeConfigInstance.Issue)
}

func GiteePushHandle(payload []byte, _ map[string]string) error {
	var raw dto.GiteePushRaw
	msgBodyErr := json.Unmarshal(payload, &raw)
	if msgBodyErr != nil {
		return msgBodyErr
	}
	rawMap := dto.StructToMap(raw)
	return handle(rawMap, config.GiteeConfigInstance.Push)
}

func GiteePrHandle(payload []byte, _ map[string]string) error {
	var raw dto.GiteePrRaw
	msgBodyErr := json.Unmarshal(payload, &raw)
	if msgBodyErr != nil {
		return msgBodyErr
	}
	rawMap := dto.StructToMap(raw)
	return handle(rawMap, config.GiteeConfigInstance.PR)
}

func GiteeNoteHandle(payload []byte, _ map[string]string) error {
	var raw dto.GiteeNoteRaw
	msgBodyErr := json.Unmarshal(payload, &raw)
	if msgBodyErr != nil {
		return msgBodyErr
	}
	rawMap := dto.StructToMap(raw)
	return handle(rawMap, config.GiteeConfigInstance.Note)
}

func EurBuildHandle(payload []byte, _ map[string]string) error {
	var raw dto.EurBuildMessageRaw
	msgBodyErr := json.Unmarshal(payload, &raw)
	if msgBodyErr != nil {
		return msgBodyErr
	}
	rawMap := dto.StructToMap(raw)
	return handle(rawMap, config.EurBuildConfigInstance.Kafka)
}

func OpenEulerMeetingHandle(payload []byte, _ map[string]string) error {
	var raw dto.OpenEulerMeetingRaw
	msgBodyErr := json.Unmarshal(payload, &raw)
	if msgBodyErr != nil {
		return msgBodyErr
	}
	rawMap := dto.StructToMap(raw)
	return handle(rawMap, config.MeetingConfigInstance.Kafka)
}
