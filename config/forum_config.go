/*
Copyright (c) Huawei Technologies Co., Ltd. 2024. All rights reserved
*/

package config

import (
	"github.com/opensourceways/message-transfer/common/kafka"
	"github.com/opensourceways/message-transfer/utils"
	"github.com/sirupsen/logrus"
)

// CveConfigInstance cve config instance.
var ForumConfigInstance ForumConfig

// ForumConfig definition of cve config.
type ForumConfig struct {
	Kafka kafka.ConsumeConfig `yaml:"kafka"`
}

// InitCVEConfig init cve config.
func InitForumConfig(configFile string) {
	cfg := new(ForumConfig)
	if err := utils.LoadFromYaml(configFile, cfg); err != nil {
		logrus.Error("Config初始化失败, err:", err)
		return
	}
	ForumConfigInstance = *cfg
}
