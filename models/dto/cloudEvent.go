package dto

import (
	"encoding/json"
	"fmt"
	cloudevents "github.com/cloudevents/sdk-go/v2"
	"github.com/opensourceways/message-transfer/common/postgresql"
	"github.com/opensourceways/message-transfer/models/do"
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
	fmt.Println(body)
	return body, err
}

func (event CloudEvents) toCloudEventDO() do.MessageCloudEventDO {
	fmt.Println(event)
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
	}
	return messageCloudEventDO
}

func (event CloudEvents) SaveDb() {
	do := event.toCloudEventDO()
	if postgresql.DB().Model(&do).Where("source=?", do.Source, "event_id = ?", do.EventId).Updates(&do).RowsAffected == 0 {
		postgresql.DB().Create(&do)
	}
}
