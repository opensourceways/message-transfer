package transfer

import (
	"encoding/json"
	"github.com/opensourceways/message-transfer/common/kafka"
	"github.com/opensourceways/message-transfer/config"
	"github.com/opensourceways/message-transfer/models/dto"
)

type GiteeHandler struct{}

func GiteeHandle(payload []byte, _ map[string]string) error {
	var raw dto.GiteeIssueRaw
	msgBodyErr := json.Unmarshal(payload, &raw)
	if msgBodyErr != nil {
		return msgBodyErr
	}
	eurBuildEvent := raw.ToCloudEventByConfig(config.GiteeConfigInstance.Issue.Topic)
	if eurBuildEvent.ID() == "" {
		return nil
	}
	kafkaSendErr := kafka.SendMsg(config.GiteeConfigInstance.Issue.Publish, &eurBuildEvent)
	if kafkaSendErr != nil {
		return kafkaSendErr
	}
	eurBuildEvent.SaveDb()
	return nil
}
