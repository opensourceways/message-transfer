package config

import (
	"github.com/opensourceways/message-transfer/common/kafka"
	"github.com/opensourceways/message-transfer/utils"
	"github.com/sirupsen/logrus"
)

// ForumConfigInstance meeting config instance.
var ForumConfigInstance ForumConfig

// ForumConfig definition of meeting config.
type ForumConfig struct {
	Kafka kafka.ConsumeConfig `yaml:"kafka"`
}

// InitForumConfig init meeting config.
func InitForumConfig(configFile string) {
	cfg := new(ForumConfig)
	if err := utils.LoadFromYaml(configFile, cfg); err != nil {
		logrus.Error("Config初始化失败, err:", err)
		return
	}
	ForumConfigInstance = *cfg
}
