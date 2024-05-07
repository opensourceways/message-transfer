package service

import (
	"github.com/IBM/sarama"
	"message-transfer/common/kafka"
	"message-transfer/service/transfer"
)

func SubscribeEurRaw() {
	cfg := kafka.ConsumeConfig{
		Topic:   "eur_build_raw",
		Address: "127.0.0.1:9092",
		Group:   "ssp_test",
		Offset:  sarama.OffsetOldest,
	}

	h := transfer.EurGroupHandler{}
	kafka.ConsumeGroup(cfg, &h)
}

func SubscribeGiteeRaw() {
	cfg := kafka.ConsumeConfig{
		Topic:   "gitee_comment_raw",
		Address: "127.0.0.1:9092",
		Group:   "ssp_test",
		Offset:  sarama.OffsetNewest,
	}

	h := transfer.GiteeGroupHandler{}
	kafka.ConsumeGroup(cfg, &h)
}
