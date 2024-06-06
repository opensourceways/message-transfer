package dto

import (
	"bytes"
	"encoding/json"
	flattener "github.com/anshal21/json-flattener"
	"github.com/araddon/dateparse"
	cloudevents "github.com/cloudevents/sdk-go/v2"
	"github.com/opensourceways/go-gitee/gitee"
	"github.com/opensourceways/message-transfer/models/bo"
	"text/template"
)

type GiteeIssueRaw struct {
	gitee.IssueEvent
}

func (raw *GiteeIssueRaw) Flatten() map[string]interface{} {
	s, _ := json.Marshal(raw)
	flatJSON, _ := flattener.FlattenJSON(string(s), flattener.DotSeparator)
	flatMap := make(map[string]interface{})
	_ = json.Unmarshal([]byte(flatJSON), &flatMap)
	return flatMap
}

/*
*
读取数据库的配置，把原始消息转换成标准的cloudevents字段
*/
func (raw *GiteeIssueRaw) ToCloudEventByConfig(sourceTopic string) CloudEvents {
	newEvent := NewCloudEvents()
	configs := bo.GetTransferConfigFromDb(sourceTopic)
	if configs != nil {
		for _, config := range configs {
			raw.transferField(&newEvent, config)
		}
		newEvent.SetData(cloudevents.ApplicationJSON, raw)
	}
	return newEvent
}

/*
*
挨个字段做映射
*/
func (raw *GiteeIssueRaw) transferField(event *CloudEvents, config bo.TransferConfig) {
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
		event.SetDataContentType(cloudevents.ApplicationCloudEventsBatchJSON)
	case "specVersion":
		event.SetSpecVersion(result)
	case "time":
		eventTime, _ := dateparse.ParseAny(result)
		event.SetTime(eventTime)
	case "user":
		event.SetExtension("user", result)
	case "sourceUrl":
		event.SetExtension("sourceurl", result)
	case "title":
		event.SetExtension("title", result)
	case "summary":
		event.SetExtension("summary", result)
	}
}
