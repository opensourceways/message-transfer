package dto

import (
	"github.com/opensourceways/go-gitee/gitee"
	"regexp"
)

type CVERaw struct {
	gitee.IssueEvent
}

func extractVariables(text string) map[string]interface{} {
	result := make(map[string]interface{})

	// 定义正则表达式来匹配每个变量
	patterns := map[string]string{
		"漏洞编号":        `漏洞编号：(.*?)\n`,
		"漏洞归属组件":      `漏洞归属组件：(.*?)\n`,
		"漏洞归属的版本":     `漏洞归属的版本：((?s).*?)\n`,
		"BaseScore":   `BaseScore：(.*?)\n`,
		"Vector":      `Vector：(.*?)\n`,
		"漏洞简述":        `漏洞简述：(.*?)漏洞公开时间`,
		"漏洞公开时间":      `漏洞公开时间：(.*?)\n`,
		"漏洞创建时间":      `漏洞创建时间：(.*?)\n`,
		"漏洞详情参考链接":    `漏洞详情参考链接：(.*?)\n`,
		"受影响版本排查":     `受影响版本排查\(受影响/不受影响\)：((?s).*?)\n\n`,
		"修复是否涉及abi变化": `修复是否涉及abi变化\(是/否\)：((?s).*?)\n\n`,
	}

	// 依次匹配并提取每个变量
	for key, pattern := range patterns {
		re := regexp.MustCompile(pattern)
		match := re.FindStringSubmatch(text)
		if len(match) > 1 {
			if key == "受影响版本排查" || key == "修复是否涉及abi变化" {
				// 将多行内容分割为数组
				lines := regexp.MustCompile(`\n`).Split(match[1], -1)
				result[key] = lines
			} else {
				result[key] = match[1]
			}
		}
	}

	return result
}
