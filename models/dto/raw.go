/*
Copyright (c) Huawei Technologies Co., Ltd. 2024. All rights reserved
*/

package dto

import (
	"bytes"
	"encoding/json"
	"reflect"
	"regexp"
	"strings"
	"text/template"
	"time"

	flattener "github.com/anshal21/json-flattener"
	"github.com/araddon/dateparse"
	cloudevents "github.com/cloudevents/sdk-go/v2"
	"github.com/sirupsen/logrus"
	"github.com/todocoder/go-stream/stream"

	"github.com/opensourceways/message-transfer/models/bo"
	"github.com/opensourceways/message-transfer/utils"
)

// Raw declare raw.
type Raw map[string]interface{}

const (
	giteeSource   = "https://gitee.com"
	meetingSource = "https://www.openEuler.org/meeting"
	cveSource     = "cve"
)

// StructToMap struct to map.
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

// Flatten flatten func.
func (raw *Raw) Flatten() map[string]interface{} {
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
func (raw *Raw) ToCloudEventByConfig(sourceTopic string) CloudEvents {
	newEvent := NewCloudEvents()
	configs := bo.GetTransferConfigFromDb(sourceTopic)
	if configs != nil {
		for _, config := range configs {
			raw.transferField(&newEvent, config)
		}
		raw.GetRelateUsers(&newEvent)
		newEvent.SetData(cloudevents.ApplicationJSON, raw)
	}
	return newEvent
}

// GetRelateUsers get relate users.
func (raw *Raw) GetRelateUsers(event *CloudEvents) {
	source := event.Source()
	if sourceGroup, ok := event.Extensions()["sourcegroup"].(string); ok {
		if result, ok := event.Extensions()["relatedusers"].(string); ok {
			lResult := strings.Split(result, ",")

			switch source {
			case giteeSource, cveSource:
				lResult = append(lResult, raw.getGiteeRelatedUsers(event, sourceGroup)...)
			case meetingSource:
				lResult = append(lResult, raw.getMeetingRelatedUsers(sourceGroup)...)
			}

			event.SetExtension("relatedusers", strings.Join(escapeCommas(raw.distinct(lResult)), ","))
		}
	}
}

func extractMentions(note string) []string {
	re := regexp.MustCompile(`@([a-zA-Z0-9_-]+)`)
	matches := re.FindAllStringSubmatch(note, -1)

	var mentions []string
	for _, match := range matches {
		if len(match) > 1 {
			mentions = append(mentions, match[1]) // match[1] 是捕获组
		}
	}
	return mentions
}

func (raw *Raw) getGiteeRelatedUsers(event *CloudEvents, sourceGroup string) []string {
	lSourceGroup := strings.Split(sourceGroup, "/")
	owner, repo := lSourceGroup[0], lSourceGroup[1]
	giteeType := event.Type()

	allAdmins, err := utils.GetAllAdmins(owner, repo)
	if err != nil {
		logrus.Errorf("get admins failed, err:%v", err)
		return []string{}
	}

	switch giteeType {
	case "pr", "push", "issue":
		return allAdmins
	case "note":
		if noteEvent, ok := (*raw)["NoteEvent"].(map[string]interface{}); ok {
			if note, ok := noteEvent["Note"].(string); ok {
				return extractMentions(note)
			} else {
				logrus.Error("Note does not exist or is not a string")
				return []string{}
			}
		} else {
			logrus.Error("NoteEvent does not exist or is not a map")
			return []string{}
		}
	default:
		return []string{}
	}
}

func (raw *Raw) getMeetingRelatedUsers(sourceGroup string) []string {
	maintainers, committers, _ := utils.GetMembersBySig(sourceGroup)
	return append(maintainers, committers...)
}

func (raw *Raw) distinct(items []string) []string {
	resultList := stream.Of(items...).Distinct(func(item string) any { return item }).ToSlice()
	var stringList []string
	for _, str := range resultList {
		if str != "" {
			stringList = append(stringList, str)
		}
	}
	return stringList
}

func escapeCommas(s []string) []string {
	var result []string
	for _, item := range s {
		if strings.HasPrefix(item, ",") {
			result = append(result, `\\`+item)
		} else {
			result = append(result, item)
		}
	}
	return result
}

/*
*
挨个字段做映射
user,sourceurl,title,summary是扩展字段
*/
func (raw *Raw) transferField(event *CloudEvents, config bo.TransferConfig) {
	tmpl := config.Template
	if config.Field == "summary" {
		logrus.Infof("the template is %v", tmpl)
	}
	parse, err := template.New("example").Funcs(
		template.FuncMap{
			"escape": func(s string) string {
				return strings.ReplaceAll(s, ",", `\\,`)
			},
		}).Parse(tmpl)
	if err != nil {
		logrus.Infof("the field is %v, the template is %v", config.Field, config.Template)
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
