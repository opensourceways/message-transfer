/*
Copyright (c) Huawei Technologies Co., Ltd. 2024. All rights reserved
*/

package dto

import (
	"bytes"
	"encoding/json"
	"reflect"
	"strings"
	"text/template"
	"time"

	flattener "github.com/anshal21/json-flattener"
	"github.com/araddon/dateparse"
	cloudevents "github.com/cloudevents/sdk-go/v2"
	"github.com/opensourceways/message-transfer/models/bo"
	"github.com/sirupsen/logrus"
)

// RawMap declare raw.
type RawMap map[string]interface{}

// StructToMap struct to map.
func StructToMap(obj interface{}) RawMap {
	objValue := reflect.ValueOf(obj)
	objType := reflect.TypeOf(obj)

	if objValue.Kind() == reflect.Ptr {
		objValue = objValue.Elem()
		objType = objType.Elem()
	}

	result := make(map[string]interface{})
	for i := 0; i < objValue.NumField(); i++ {
		field := objValue.Field(i)
		fieldType := objType.Field(i)
		fieldName := fieldType.Name

		// 解引用指针字段
		for field.Kind() == reflect.Ptr {
			if field.IsNil() {
				result[fieldName] = nil
				break
			}
			field = field.Elem()
		}

		if field.Kind() == reflect.Invalid {
			continue
		}

		switch field.Kind() {
		case reflect.Struct:
			if field.Type() == reflect.TypeOf(time.Time{}) {
				result[fieldName] = field.Interface().(time.Time).Format(time.RFC3339)
			} else {
				result[fieldName] = StructToMap(field.Interface())
			}
		default:
			result[fieldName] = field.Interface()
		}
	}
	return result
}

func (raw *RawMap) Flatten() map[string]interface{} {
	s, err := json.Marshal(raw)
	if err != nil {
		logrus.Error(err)
		return nil
	}
	flatJSON, _ := flattener.FlattenJSON(string(s), flattener.DotSeparator)
	flatMap := make(map[string]interface{})
	err = json.Unmarshal([]byte(flatJSON), &flatMap)
	if err != nil {
		logrus.Error(err)
		return nil
	}
	return flatMap
}

//ToCloudEventByConfig
/*
*
读取数据库的配置，把原始消息转换成标准的cloudevents字段
*/
func (raw *RawMap) ToCloudEventByConfig(sourceTopic string) CloudEvents {
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
user,sourceurl,title,summary是扩展字段
*/
func (raw *RawMap) transferField(event *CloudEvents, config bo.TransferConfig) {
	tmpl := config.Template
	parse, err := template.New("example").Funcs(
		template.FuncMap{
			"escape": func(s string) string {
				return strings.ReplaceAll(s, ",", `\\,`)
			},
		}).Parse(tmpl)
	if err != nil {
		logrus.Error(config.Field, err)
	}
	t := template.Must(parse, nil)
	var resultBuffer bytes.Buffer
	err = t.Execute(&resultBuffer, raw)
	if err != nil {
		logrus.Error(err.Error())
		return
	}
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
	case "specVersion":
		event.SetSpecVersion(result)
	case "time":
		eventTime, _ := dateparse.ParseAny(result)
		event.SetTime(eventTime)
	case "user":
		event.SetExtension("user", result)
	case "sourceUrl":
		event.SetExtension("sourceurl", result)
	case "sourceGroup":
		event.SetExtension("sourcegroup", result)
	case "title":
		event.SetExtension("title", result)
	case "summary":
		event.SetExtension("summary", result)
	case "relatedUsers":
		event.SetExtension("relatedusers", result)
	case "mailTitle":
		event.SetExtension("mailtitle", result)
	case "mailSummary":
		event.SetExtension("mailsummary", result)
	}
}
