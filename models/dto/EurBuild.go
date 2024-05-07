package dto

import (
	"bytes"
	"encoding/json"
	flattener "github.com/anshal21/json-flattener"
	cloudevents "github.com/cloudevents/sdk-go/v2"
	"github.com/cloudevents/sdk-go/v2/event"
	"message-transfer/models/bo"
	"text/template"
	"time"
)

type EurBuildMessageRaw struct {
	Properties struct {
		AppID           interface{} `json:"app_id"`
		ClusterID       interface{} `json:"cluster_id"`
		ContentEncoding string      `json:"content_encoding"`
		ContentType     string      `json:"content_type"`
		CorrelationID   interface{} `json:"correlation_id"`
		DeliveryMode    int         `json:"delivery_mode"`
		Expiration      interface{} `json:"expiration"`
		Headers         struct {
			FedoraMessagingSchema      string `json:"fedora_messaging_schema"`
			FedoraMessagingSeverity    int    `json:"fedora_messaging_severity"`
			FedoraMessagingUserKkleine bool   `json:"fedora_messaging_user_kkleine"`
			Priority                   int    `json:"priority"`
			SentAt                     string `json:"sent-at"`
			XReceivedFrom              []struct {
				ClusterName string `json:"cluster-name"`
				Exchange    string `json:"exchange"`
				Redelivered bool   `json:"redelivered"`
				URI         string `json:"uri"`
			} `json:"x-received-from"`
		} `json:"headers"`
		MessageID string      `json:"message_id"`
		Priority  interface{} `json:"priority"`
		ReplyTo   interface{} `json:"reply_to"`
		Timestamp interface{} `json:"timestamp"`
		Type      interface{} `json:"type"`
		UserID    interface{} `json:"user_id"`
	} `json:"_properties"`
	Body struct {
		Build   int    `json:"build"`
		Chroot  string `json:"chroot"`
		Copr    string `json:"copr"`
		IP      string `json:"ip"`
		Owner   string `json:"owner"`
		PID     int    `json:"pid"`
		Pkg     string `json:"pkg"`
		Status  int    `json:"status"`
		User    string `json:"user"`
		Version string `json:"version"`
		What    string `json:"what"`
		Who     string `json:"who"`
	} `json:"body"`
	Queue    string `json:"queue"`
	Severity int    `json:"severity"`
	Topic    string `json:"topic"`
}

func (raw *EurBuildMessageRaw) Flatten() map[string]interface{} {
	s, _ := json.Marshal(raw)
	flatJSON, _ := flattener.FlattenJSON(string(s), flattener.DotSeparator)
	flatMap := make(map[string]interface{})
	_ = json.Unmarshal([]byte(flatJSON), &flatMap)
	return flatMap
}

func (raw *EurBuildMessageRaw) ToCloudEventByConfig(sourceTopic string) CloudEvents {
	newEvent := cloudevents.NewEvent()
	configs := bo.GetTransferConfigFromDb(sourceTopic)
	for _, config := range configs {
		raw.transferField(newEvent, config)
	}
	newEvent.SetData(cloudevents.ApplicationJSON, raw)
	return CloudEvents{newEvent}

}

func (raw *EurBuildMessageRaw) transferField(event event.Event, config bo.TransferConfig) {
	tmpl := config.Template
	t := template.Must(template.New("example").Parse(tmpl))
	var resultBuffer bytes.Buffer
	t.Execute(&resultBuffer, raw)
	result := resultBuffer.String()
	switch config.Field {
	case "id":
		event.SetID(result)
	case "source":
		event.SetSource(result)
	case "dataSchema":
		event.SetDataSchema(result)
	case "type":
		event.SetType(result)
	case "dataContentType":
		event.SetDataContentType(result)
	case "specVersion":
		event.SetSpecVersion(result)
	case "time":
		eventTime, _ := time.Parse(time.RFC3339, result)
		event.SetTime(eventTime)
	}
}
