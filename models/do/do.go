package do

import (
	"time"

	"github.com/opensourceways/message-transfer/common/postgresql"
	"gorm.io/datatypes"
)

type JSONB []interface{}

type MessageCloudEventDO struct {
	postgresql.CommonModel
	EventId         string         `gorm:"column:event_id"`
	Source          string         `gorm:"column:source"`
	EventType       string         `gorm:"column:type"`
	DataContentType string         `gorm:"column:data_content_type"`
	DataSchema      string         `gorm:"column:data_schema"`
	SpecVersion     string         `gorm:"column:spec_version"`
	Time            time.Time      `gorm:"column:time"`
	User            string         `gorm:"column:user"`
	SourceUrl       string         `gorm:"column:source_url"`
	DataJson        datatypes.JSON `gorm:"column:data_json"`
	Title           string         `gorm:"column:title"`
	Summary         string         `gorm:"column:summary"`
	SourceGroup     string         `gorm:"column:source_group"`
	RelatedUsers    string         `gorm:"column:related_users"`
	MailTitle       string         `gorm:"column:mail_title"`
	MailSummary     string         `gorm:"column:mail_summary"`
}

func (m *MessageCloudEventDO) TableName() string {
	return "message_center.cloud_event_message"
}

type SubScribeConfigDO struct {
	postgresql.CommonModel
	Source     string         `gorm:"column:source"`
	EventType  string         `gorm:"column:type"`
	Version    string         `gorm:"column:version"`
	ModeFilter datatypes.JSON `gorm:"column:mod_filter"`
	IsDeleted  bool           `gorm:"column:is_deleted"`
}

func (m *SubScribeConfigDO) TableName() string {
	return "subscribe_config"
}

type PushConfigDO struct {
	postgresql.CommonModel
	SubScribeId int    `gorm:"column:subscribe_id"`
	PushType    string `gorm:"column:type"`
	PushAddress string `gorm:"column:version"`
	IsDeleted   bool   `gorm:"column:is_deleted"`
}

func (m *SubScribeConfigDO) PushConfigDO() string {
	return "push_config"
}
