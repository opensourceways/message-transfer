/*
Copyright (c) Huawei Technologies Co., Ltd. 2024. All rights reserved
*/

package config

import (
	"github.com/sirupsen/logrus"

	"github.com/opensourceways/message-transfer/common/kafka"
	"github.com/opensourceways/message-transfer/utils"
)

// CveConfigInstance cve config instance.
var CveConfigInstance CVEConfig

// CVEConfig definition of cve config.
type CVEConfig struct {
	Kafka kafka.ConsumeConfig `yaml:"kafka"`
}

// InitCVEConfig init cve config.
func InitCVEConfig(configFile string) {
	cfg := new(CVEConfig)
	if err := utils.LoadFromYaml(configFile, cfg); err != nil {
		logrus.Error("Config初始化失败, err:", err)
		return
	}
	CveConfigInstance = *cfg
}
