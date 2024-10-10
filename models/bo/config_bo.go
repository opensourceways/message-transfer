/*
Copyright (c) Huawei Technologies Co., Ltd. 2024. All rights reserved
*/

// Package bo transfer config.
package bo

import (
	"github.com/opensourceways/message-transfer/common/postgresql"
)

// TransferConfig transfer config.
type TransferConfig struct {
	SourceTopic string `gorm:"column:source_topic"`
	Field       string `gorm:"column:field"`
	Template    string `gorm:"column:template"`
}

// GetTransferConfigFromDb get transfer config from db.
func GetTransferConfigFromDb(sourceTopic string) []TransferConfig {
	var transferConfigs []TransferConfig
	postgresql.
		DB().
		Table("message_center.transfer_config").
		Where("source_topic=?", sourceTopic).
		Where("is_deleted=?", false).
		Find(&transferConfigs)
	return transferConfigs
}
