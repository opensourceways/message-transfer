package dto

import (
	"github.com/araddon/dateparse"
	"github.com/opensourceways/message-transfer/models/bo"
	"github.com/stretchr/testify/assert"
	"testing"
)

// 模拟 CloudEvents

// 模拟 TransferConfig

func TestTransferField(t *testing.T) {
	tests := []struct {
		name           string
		raw            Raw
		config         bo.TransferConfig
		expectedField  string
		expectedResult string
	}{
		{
			name: "Test ID field",
			raw: Raw{
				"id": "12345",
			},
			config: bo.TransferConfig{
				Field:    "id",
				Template: `{{.id}}-test`,
			},
			expectedField:  "id",
			expectedResult: "12345-test",
		},
		{
			name: "Test Source field",
			raw: Raw{
				"source": "source-test",
			},
			config: bo.TransferConfig{
				Field:    "source",
				Template: `{{.source}}-example`,
			},
			expectedField:  "source",
			expectedResult: "source-test-example",
		},
		{
			name: "Test Time field",
			raw: Raw{
				"time": "2023-09-15T14:12:00Z",
			},
			config: bo.TransferConfig{
				Field:    "time",
				Template: `{{.time}}`,
			},
			expectedField:  "time",
			expectedResult: "2023-09-15T14:12:00Z",
		},
		{
			name: "Test User extension",
			raw: Raw{
				"user": "user123",
			},
			config: bo.TransferConfig{
				Field:    "user",
				Template: `{{.user}}-user`,
			},
			expectedField:  "user",
			expectedResult: "user123-user",
		},
		// 更多测试用例可以根据你需要的字段类型扩展
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// 创建 CloudEvent 实例
			event := CloudEvents{}

			// 执行 transferField
			raw := Raw(tt.raw) // 使用传入的原始数据
			raw.transferField(&event, tt.config)

			// 验证不同的字段
			switch tt.config.Field {
			case "id":
				assert.Equal(t, tt.expectedResult, event.ID())
			case "source":
				assert.Equal(t, tt.expectedResult, event.Source())
			case "time":
				parsedTime, _ := dateparse.ParseAny(tt.expectedResult)
				assert.Equal(t, parsedTime, event.Time())
			case "user":
				assert.Equal(t, tt.expectedResult, event.Extensions()["user"])
				// 根据你需要的字段继续添加更多断言
			}
		})
	}
}
