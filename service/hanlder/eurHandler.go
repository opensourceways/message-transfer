package hanlder

import (
	"encoding/json"
	"fmt"
	"github.com/IBM/sarama"
	"github.com/opensourceways/kafka-lib/mq"
	"message-transfer/common/postgresql"
	"message-transfer/models/event"
	"message-transfer/models/messageadapter"
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
		var msg mq.Message
		err := json.Unmarshal(message.Value, &msg)
		if err != nil {
			return err
		}
		var raw event.EurBuildRaw

		msgBodyErr := json.Unmarshal(msg.Body, &raw)
		if msgBodyErr != nil {
			return err
		}
		fmt.Printf("Received message with offset %d: %s\n", message.Offset, raw)

		transferErr := publishEurEvent(raw)
		if transferErr != nil {
			return transferErr
		}
		save(raw)
		session.MarkMessage(message, "")
	}
	return nil
}

func publishEurEvent(raw event.EurBuildRaw) error {
	eurBuildEvent := raw.ToCloudEvent()
	sendErr := messageadapter.SendMsg("eur_build_event", &eurBuildEvent)
	if sendErr != nil {
		return sendErr
	}
	return nil
}

func save(raw event.EurBuildRaw) {
	do := raw.ToCloudEventDO()
	res := postgresql.DB().Table("message_center.cloud_event_message").Create(&do)
	fmt.Println(res)
}
