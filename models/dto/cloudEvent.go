package dto

import (
	"encoding/json"
	"fmt"

	cloudevents "github.com/cloudevents/sdk-go/v2"
	"github.com/opensourceways/message-transfer/common/postgresql"
	"github.com/opensourceways/message-transfer/models/do"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm/clause"
)

type CloudEvents struct {
	cloudevents.Event
}

func NewCloudEvents() CloudEvents {
	return CloudEvents{
		Event: cloudevents.NewEvent(cloudevents.VersionV1),
	}
}

func (event CloudEvents) Message() ([]byte, error) {
	body, err := json.Marshal(event)
	logrus.Info(body)
	return body, err
}

func (event CloudEvents) SaveDb() {
	eventDO := event.toCloudEventDO()
	res := postgresql.DB().Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "event_id"}, {Name: "source"}},
		DoUpdates: clause.AssignmentColumns([]string{"event_id", "source_url", "source_group", "summary", "data_schema", "data_content_type", "spec_version", "time", "user", "data_json", "title", "related_users"}),
	}).Create(&eventDO)
	logrus.Info(res)
	logrus.Info("插入成功")
}

func (event CloudEvents) toCloudEventDO() do.MessageCloudEventDO {
	fmt.Println(event)
	//relatedUsers := strings.Split(event.Extensions()["relatedusers"].(string), ",")
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
	}
	return messageCloudEventDO
}
