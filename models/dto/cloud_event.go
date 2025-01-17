/*
Copyright (c) Huawei Technologies Co., Ltd. 2024. All rights reserved
*/

// Package dto models dto.
package dto

import (
	"encoding/json"

	cloudevents "github.com/cloudevents/sdk-go/v2"
	"github.com/sirupsen/logrus"
	"golang.org/x/xerrors"
	"gorm.io/gorm/clause"

	"github.com/opensourceways/message-transfer/common/postgresql"
	"github.com/opensourceways/message-transfer/models/do"
)

// CloudEvents cloud event.
type CloudEvents struct {
	cloudevents.Event
}

// NewCloudEvents create a new cloud event.
func NewCloudEvents() CloudEvents {
	return CloudEvents{
		Event: cloudevents.NewEvent(cloudevents.VersionV1),
	}
}

// Message marshal event to msg.
func (event CloudEvents) Message() ([]byte, error) {
	body, err := json.Marshal(event)
	return body, err
}

// SaveDb save cloud event into db.
func (event CloudEvents) SaveDb() error {
	eventDO := event.toCloudEventDO()
	result := postgresql.DB().Clauses(clause.OnConflict{
		Columns: []clause.Column{{Name: "event_id"}},
		DoUpdates: clause.AssignmentColumns([]string{"event_id", "source_url", "source_group",
			"summary", "data_schema", "data_content_type", "spec_version", "time", "user",
			"data_json", "title", "related_users", "mail_title", "mail_summary"}),
	}).Create(&eventDO)
	if result.Error != nil {
		return xerrors.Errorf("save DB failed, the err: %v", result.Error)
	}
	logrus.Info("插入成功")
	return nil
}

func (event CloudEvents) toCloudEventDO() do.MessageCloudEventDO {
	if event.Extensions()["sourcegroup"].(string) == "openeuler/infrastructure" {
		logrus.Errorf("the saveDB result is %v", event.Extensions()["relatedusers"].(string))
	}
	messageCloudEventDO := do.MessageCloudEventDO{
		Source:          event.Source(),
		Time:            event.Time(),
		EventType:       event.Type(),
		SpecVersion:     event.SpecVersion(),
		DataSchema:      event.DataSchema(),
		DataContentType: event.DataContentType(),
		EventId:         event.ID(),
		DataJson:        event.Data(),
		User:            event.Extensions()["user"].(string),
		SourceUrl:       event.Extensions()["sourceurl"].(string),
		Title:           event.Extensions()["title"].(string),
		Summary:         event.Extensions()["summary"].(string),
		SourceGroup:     event.Extensions()["sourcegroup"].(string),
		RelatedUsers:    "{" + event.Extensions()["relatedusers"].(string) + "}",
		MailTitle:       event.Extensions()["mailtitle"].(string),
		MailSummary:     event.Extensions()["mailsummary"].(string),
	}
	return messageCloudEventDO
}
