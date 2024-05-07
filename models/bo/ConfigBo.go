package bo

import (
	"message-transfer/common/postgresql"
)

type TransferConfig struct {
	SourceTopic string `gorm:"column:source_topic"`
	Field       string `gorm:"column:field"`
	Template    string `gorm:"column:template"`
}

func GetTransferConfigFromDb(sourceTopic string) []TransferConfig {
	var transferConfigs []TransferConfig
	postgresql.DB().Table("message_center.transfer_config").Find(&transferConfigs, "source_topic=?", sourceTopic, "is_deleted", false)
	return transferConfigs
}
