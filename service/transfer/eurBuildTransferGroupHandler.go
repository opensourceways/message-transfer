package transfer

import (
	"encoding/json"
	"github.com/IBM/sarama"
	"github.com/opensourceways/message-transfer/common/kafka"
	"github.com/opensourceways/message-transfer/config"
	"github.com/opensourceways/message-transfer/models/dto"
)

type EurHandler struct{}

func Handle(payload []byte, _ map[string]string) error {
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

type EurGroupHandler struct{}

func (h EurGroupHandler) Setup(_ sarama.ConsumerGroupSession) error {
	return nil
}

func (h EurGroupHandler) Cleanup(_ sarama.ConsumerGroupSession) error {
	return nil
}

func (h EurGroupHandler) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for message := range claim.Messages() {
		var raw dto.EurBuildMessageRaw
		msgBodyErr := json.Unmarshal(message.Value, &raw)
		if msgBodyErr != nil {
			return msgBodyErr
		}
		eurBuildEvent := raw.ToCloudEventByConfig(message.Topic)
		if eurBuildEvent.ID() == "" {
			session.MarkMessage(message, "")
			continue
		}
		kafkaSendErr := kafka.SendMsg(config.EurBuildConfigInstance.Kafka.Publish, &eurBuildEvent)
		if kafkaSendErr != nil {
			return kafkaSendErr
		}
		eurBuildEvent.SaveDb()
		session.MarkMessage(message, "")
	}
	return nil
}
