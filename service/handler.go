/*
Copyright (c) Huawei Technologies Co., Ltd. 2024. All rights reserved
*/

// Package service handle func.
package service

import (
	"time"

	"github.com/sirupsen/logrus"

	"github.com/opensourceways/message-transfer/common/kafka"
	"github.com/opensourceways/message-transfer/models/dto"
	"github.com/opensourceways/message-transfer/utils"
)

func handle(raw dto.RawMap, cfg kafka.ConsumeConfig) error {
	time.Sleep(utils.GetConsumeSleepTime())
	event := raw.ToCloudEventByConfig(cfg.Topic)
	if event.ID() == "" {
		return nil
	}
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
