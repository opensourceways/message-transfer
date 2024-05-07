package dto

import (
	"encoding/json"
	cloudevents "github.com/cloudevents/sdk-go/v2"
	"message-transfer/common/postgresql"
	"message-transfer/models/do"
)

type CloudEvents struct {
	cloudevents.Event
}

func (event CloudEvents) Message() ([]byte, error) {
	return json.Marshal(event)
}

func (event CloudEvents) toCloudEventDO() do.MessageCloudEventDO {
	messageCloudEventDO := do.MessageCloudEventDO{
		Source:          event.Source(),
		Time:            event.Time(),
		EventType:       event.Type(),
		SpecVersion:     event.SpecVersion(),
		DataSchema:      event.DataSchema(),
		DataContentType: event.DataContentType(),
		EventId:         event.ID(),
		DataJson:        event.Data(),
	}
	return messageCloudEventDO
}

func (event CloudEvents) SaveDb() {
	do := event.toCloudEventDO()
	if postgresql.DB().Model(&do).Where("source=?", do.Source, "event_id = ?", do.EventId).Updates(&do).RowsAffected == 0 {
		postgresql.DB().Create(&do)
	}
}
