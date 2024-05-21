package transfer

import (
	"encoding/json"
	"fmt"
	"github.com/IBM/sarama"
	"message-transfer/common/kafka"
	"message-transfer/models/dto"
)

type EurHandler struct{}

func (eurHandler *EurHandler) handle(message []byte) error {
	fmt.Println(message)
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
		kafkaSendErr := kafka.SendMsg("eur_build_event", &eurBuildEvent)
		if kafkaSendErr != nil {
			return kafkaSendErr
		}
		eurBuildEvent.SaveDb()
		session.MarkMessage(message, "")
	}
	return nil
}
