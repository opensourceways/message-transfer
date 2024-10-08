/*
Copyright (c) Huawei Technologies Co., Ltd. 2024. All rights reserved
*/

// Package config provides functionality for managing application configuration.
package config

import (
	"github.com/opensourceways/message-transfer/common/kafka"
	"github.com/opensourceways/message-transfer/common/postgresql"
	"github.com/opensourceways/message-transfer/utils"
)

// Config is a struct that represents the overall configuration for the application.
type Config struct {
	Kafka      kafka.Config      `json:"kafka"`
	Postgresql postgresql.Config `json:"postgresql"`
	Utils      utils.Config      `json:"utils"`
}
