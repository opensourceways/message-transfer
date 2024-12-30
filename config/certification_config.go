package config

import (
	"github.com/sirupsen/logrus"

	"github.com/opensourceways/message-transfer/common/kafka"
	"github.com/opensourceways/message-transfer/utils"
)

// CertificationConfigInstance meeting config instance.
var CertificationConfigInstance CertificationConfig

// CertificationConfig definition of meeting config.
type CertificationConfig struct {
	Kafka kafka.ConsumeConfig `yaml:"kafka"`
}

// InitCertificationConfig init meeting config.
func InitCertificationConfig(configFile string) {
	cfg := new(CertificationConfig)
	if err := utils.LoadFromYaml(configFile, cfg); err != nil {
		logrus.Error("Config初始化失败, err:", err)
		return
	}
	CertificationConfigInstance = *cfg
}
