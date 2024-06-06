package transfer

import (
	"encoding/json"
	"github.com/opensourceways/message-transfer/common/kafka"
	"github.com/opensourceways/message-transfer/config"
	"github.com/opensourceways/message-transfer/models/dto"
)

type EurHandler struct{}

func EurBuildHandle(payload []byte, _ map[string]string) error {
	var raw dto.EurBuildMessageRaw
	msgBodyErr := json.Unmarshal(payload, &raw)
	if msgBodyErr != nil {
		return msgBodyErr
	}
	eurBuildEvent := raw.ToCloudEventByConfig(config.EurBuildConfigInstance.Kafka.Topic)
	if eurBuildEvent.ID() == "" {
		return nil
	}
	kafkaSendErr := kafka.SendMsg(config.EurBuildConfigInstance.Kafka.Publish, &eurBuildEvent)
	if kafkaSendErr != nil {
		return kafkaSendErr
	}
	eurBuildEvent.SaveDb()
	return nil
}
