package service

import (
	"encoding/json"

	"github.com/sirupsen/logrus"

	"github.com/opensourceways/message-transfer/common/kafka"
	"github.com/opensourceways/message-transfer/config"
	"github.com/opensourceways/message-transfer/models/dto"
	"github.com/opensourceways/message-transfer/models/dto/gitee"
)

func handleEvent(rawEvent dto.RawEvent, cfg kafka.ConsumeConfig) error {
	event := rawEvent.ToCloudEventsByConfig(cfg.Topic)
	if event.ID() == "" {
		logrus.Errorf("event id is empty")
		return nil
	}
	rawEvent.GetTodoUsers(event)
	rawEvent.GetRelateUsers(event)
	rawEvent.GetFollowUsers(event)
	rawEvent.IsDone(event)
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
	var raw gitee.IssueRaw
	msgBodyErr := json.Unmarshal(payload, &raw)
	if msgBodyErr != nil {
		return msgBodyErr
	}
	return handleEvent(&raw, config.GiteeConfigInstance.Issue)
}

// GiteePrHandle handle gitee pr raw.
func GiteePrHandle(payload []byte, _ map[string]string) error {
	var raw gitee.PrRaw
	msgBodyErr := json.Unmarshal(payload, &raw)
	if msgBodyErr != nil {
		return msgBodyErr
	}
	return handleEvent(&raw, config.GiteeConfigInstance.PR)
}

// GiteeNoteHandle handle gitee note raw.
func GiteeNoteHandle(payload []byte, _ map[string]string) error {
	var raw gitee.NoteRaw
	msgBodyErr := json.Unmarshal(payload, &raw)
	if msgBodyErr != nil {
		return msgBodyErr
	}
	return handleEvent(&raw, config.GiteeConfigInstance.Note)
}

// CVEHandle handle cve issue raw.
func CVEHandle(payload []byte, _ map[string]string) error {
	var raw dto.CVEIssueRaw
	msgBodyErr := json.Unmarshal(payload, &raw)
	if msgBodyErr != nil {
		return msgBodyErr
	}
	return handleEvent(&raw, config.CveConfigInstance.Kafka)
}

// EurBuildHandle handle eur build raw.
func EurBuildHandle(payload []byte, _ map[string]string) error {
	var raw dto.EurBuildMessageRaw
	msgBodyErr := json.Unmarshal(payload, &raw)
	if msgBodyErr != nil {
		return msgBodyErr
	}
	return handleEvent(&raw, config.EurBuildConfigInstance.Kafka)
}

// OpenEulerMeetingHandle handle openEuler meeting raw.
func OpenEulerMeetingHandle(payload []byte, _ map[string]string) error {
	var raw dto.OpenEulerMeetingRaw
	msgBodyErr := json.Unmarshal(payload, &raw)
	logrus.Infof("receive meeting message, Title:%v", raw.Msg.Topic)
	if msgBodyErr != nil {
		logrus.Errorf("unmarshal meeting message failed, err:%v", msgBodyErr)
		return msgBodyErr
	}
	return handleEvent(&raw, config.MeetingConfigInstance.Kafka)
}

// OpenEulerForumHandle handle openEuler forum raw.
func OpenEulerForumHandle(payload []byte, _ map[string]string) error {
	var raw dto.Notification
	msgBodyErr := json.Unmarshal(payload, &raw)
	if msgBodyErr != nil {
		logrus.Errorf("unmarshal forum message failed, err:%v", msgBodyErr)
		return msgBodyErr
	}
	return handleEvent(&raw, config.ForumConfigInstance.Kafka)
}

// CertificationHandle handle certification raw.
func CertificationHandle(payload []byte, _ map[string]string) error {
	var raw dto.CertificationRaw
	msgBodyErr := json.Unmarshal(payload, &raw)
	if msgBodyErr != nil {
		logrus.Errorf("unmarshal certification message failed, err:%v", msgBodyErr)
		return msgBodyErr
	}
	return handleEvent(&raw, config.CertificationConfigInstance.Kafka)
}
