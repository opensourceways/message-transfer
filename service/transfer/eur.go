package transfer

import (
	"github.com/IBM/sarama"
	"message-transfer/models/messageadapter"
	"message-transfer/service/hanlder"
)

func ConsumeEur() {
	cfg := messageadapter.ConsumeConfig{
		Topic:   "eur_build_raw",
		Address: "0.0.0.0:9092",
		Group:   "ssp_test",
		Offset:  sarama.OffsetOldest,
	}

	h := hanlder.EurGroupHandler{}
	messageadapter.ConsumeGroup(cfg, &h)
}
