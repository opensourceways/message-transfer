/*
Copyright (c) Huawei Technologies Co., Ltd. 2024. All rights reserved
*/

package config

import (
	"github.com/opensourceways/message-transfer/common/kafka"
	"github.com/opensourceways/message-transfer/utils"
	"github.com/sirupsen/logrus"
)

// GiteeConfigInstance gitee config instance.
var GiteeConfigInstance GiteeConfig

// GiteeConfig definition of gitee config.
type GiteeConfig struct {
	Issue kafka.ConsumeConfig `yaml:"issue"`
	Push  kafka.ConsumeConfig `yaml:"push"`
	PR    kafka.ConsumeConfig `yaml:"pr"`
	Note  kafka.ConsumeConfig `yaml:"note"`
}

// InitGiteeConfig init gitee config.
func InitGiteeConfig(configFile string) {
	cfg := new(GiteeConfig)
	if err := utils.LoadFromYaml(configFile, cfg); err != nil {
		logrus.Error("Config初始化失败, err:", err)
		return
	}
	GiteeConfigInstance = *cfg
}
