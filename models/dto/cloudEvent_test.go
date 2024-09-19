package dto

import (
	"testing"
	"time"

	"github.com/opensourceways/message-transfer/models/do"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

// Mock DB (可以使用 sqlmock 或者简单的模拟函数来替代实际的数据库操作)
var mockDB *gorm.DB

// 模拟 postgresql.DB() 返回值
func mockDBFunc() *gorm.DB {
	return mockDB
}

func TestCloudEvents_Message(t *testing.T) {
	event := NewCloudEvents()
	event.SetSource("test-source")
	event.SetID("12345")
	event.SetType("test-type")

	// 测试 Message 方法
	body, err := event.Message()

	// 检查没有错误
	assert.Nil(t, err)

	// 检查 JSON 序列化后的内容
	expected := `{"specversion":"1.0","id":"12345","source":"test-source","type":"test-type"}`
	assert.JSONEq(t, expected, string(body))
}

func TestCloudEvents_SaveDb(t *testing.T) {
	event := NewCloudEvents()
	event.SetSource("test-source")
	event.SetID("12345")
	event.SetType("test-type")
	event.SetDataSchema("http://example.com/schema")
	event.SetDataContentType("application/json")
	event.SetTime(time.Now())
	event.SetExtension("user", "test-user")
	event.SetExtension("sourceurl", "http://example.com/source")
	event.SetExtension("title", "Test Title")
	event.SetExtension("summary", "Test Summary")
	event.SetExtension("sourcegroup", "Test Group")
	event.SetExtension("relatedusers", "user1,user2")

	// 模拟 toCloudEventDO 返回值
	expectedDO := do.MessageCloudEventDO{
		Source:          "test-source",
		Time:            event.Time(),
		EventType:       "test-type",
		SpecVersion:     "1.0",
		DataSchema:      "http://example.com/schema",
		DataContentType: "application/json",
		EventId:         "12345",
		User:            "test-user",
		SourceUrl:       "http://example.com/source",
		Title:           "Test Title",
		Summary:         "Test Summary",
		SourceGroup:     "Test Group",
		RelatedUsers:    "{user1,user2}",
	}

	// 模拟数据库插入结果
	mockDB = &gorm.DB{} // 使用 mock 数据库

	// 替换 postgresql.DB() 方法
	postgresql.DB = mockDBFunc

	// 调用 SaveDb 方法
	err := event.SaveDb()

	// 检查没有错误
	assert.Nil(t, err)

	// 检查生成的 DO 对象是否正确
	actualDO := event.toCloudEventDO()
	assert.Equal(t, expectedDO, actualDO)

	// 检查是否正确调用了数据库插入
	// (可以使用 sqlmock 或其他 mock 工具进行数据库行为的验证)
}
