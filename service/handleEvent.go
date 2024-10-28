package service

import (
	"encoding/json"
	"github.com/opensourceways/message-transfer/common/kafka"
	"github.com/opensourceways/message-transfer/config"
	"github.com/opensourceways/message-transfer/models/dto"
	"github.com/opensourceways/message-transfer/models/dto/gitee"
	"github.com/sirupsen/logrus"
)

func handleEvent(rawEvent dto.RawEvent, cfg kafka.ConsumeConfig) error {
	event := rawEvent.ToCloudEventsByConfig()
	if event.ID() == "" {
		return nil
	}
	rawEvent.GetTodoUsers(event)
	rawEvent.GetRelateUsers(event)
	rawEvent.GetFollowUsers(event)
	err := event.SaveDb()
	if err != nil {
		logrus.Errorf("saveDb failed, err:%v", err)
		return err
	}
	kafkaSendErr := kafka.SendMsg(cfg.Publish, &event)
	if kafkaSendErr != nil {
		return kafkaSendErr
	}
	return nil
}

// GiteeIssueHandle handle gitee issue raw.
func GiteeIssueHandle(payload []byte, _ map[string]string) error {
	var raw gitee.GiteeIssueRaw
	msgBodyErr := json.Unmarshal(payload, &raw)
	if msgBodyErr != nil {
		return msgBodyErr
	}
	return handleEvent(&raw, config.GiteeConfigInstance.Issue)
}

func EurBuildHandle(payload []byte, _ map[string]string) error {
	var raw dto.EurBuildMessageRaw
	msgBodyErr := json.Unmarshal(payload, &raw)
	if msgBodyErr != nil {
		return msgBodyErr
	}
	return handleEvent(&raw, config.EurBuildConfigInstance.Kafka)
}
