package service

import (
	"github.com/IBM/sarama"
	"github.com/opensourceways/message-transfer/common/kafka"
	"github.com/opensourceways/message-transfer/service/transfer"
)

func SubscribeEurRaw() {
	cfg := kafka.ConsumeConfig{
		Topic:   "eur_build_raw",
		Address: "182.160.6.195:9094",
		Group:   "message-transfer",
		Offset:  sarama.OffsetOldest,
	}

	h := transfer.EurGroupHandler{}
	kafka.ConsumeGroup(cfg, &h)
}
