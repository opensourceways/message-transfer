package dto

import (
	"bytes"
	"encoding/json"
	"fmt"
	"reflect"
	"text/template"
	"time"

	flattener "github.com/anshal21/json-flattener"
	"github.com/araddon/dateparse"
	cloudevents "github.com/cloudevents/sdk-go/v2"
	"github.com/opensourceways/message-transfer/models/bo"
	"github.com/sirupsen/logrus"
)

type Raw map[string]interface{}

func StructToMap(obj interface{}) map[string]interface{} {
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

func ToMap(in interface{}) (Raw, error) {
	out := make(map[string]interface{})
	v := reflect.ValueOf(in)

	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	if v.Kind() != reflect.Struct {
		return nil, fmt.Errorf("ToMap only accepts struct or struct pointer; got %T", v)
	}

	t := v.Type()

	for i := 0; i < v.NumField(); i++ {
		fi := t.Field(i)
		out[fi.Name] = v.Field(i).Interface()
	}

	return out, nil
}

func (raw *Raw) Flatten() map[string]interface{} {
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
func (raw *Raw) ToCloudEventByConfig(sourceTopic string) CloudEvents {
	newEvent := NewCloudEvents()
	configs := bo.GetTransferConfigFromDb(sourceTopic)
	if configs != nil {
		for _, config := range configs {
			raw.transferField(&newEvent, config)
		}

		logrus.Infof("the event source is %v", newEvent.Source())
		logrus.Infof("the event sourgroup is %v", newEvent.Extensions()["sourcegroup"].(string))

		newEvent.SetData(cloudevents.ApplicationJSON, raw)
	}
	return newEvent
}

/*
*
挨个字段做映射
user,sourceurl,title,summary是扩展字段
*/
func (raw *Raw) transferField(event *CloudEvents, config bo.TransferConfig) {
	tmpl := config.Template
	logrus.Infof("the tmpl is %v", tmpl)
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
	}
}
