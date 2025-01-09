/*
Copyright (c) Huawei Technologies Co., Ltd. 2024. All rights reserved
*/

// Package dto models dto of cve
package dto

import (
	"regexp"
	"slices"
	"strconv"
	"strings"

	"github.com/opensourceways/go-gitee/gitee"

	"github.com/opensourceways/message-transfer/utils"
)

// CVEIssueRaw cve issue raw.
type CVEIssueRaw struct {
	gitee.IssueEvent
}

func (cveIssueRaw *CVEIssueRaw) ToCloudEventsByConfig(topic string) CloudEvents {
	rawMap := cveIssueRaw.ToMap()
	return rawMap.ToCloudEventByConfig(topic)
}

func (cveIssueRaw *CVEIssueRaw) GetRelateUsers(events CloudEvents) {
	events.SetExtension("relatedusers", "")
}

func (cveIssueRaw *CVEIssueRaw) followUsers() []string {
	sigGroup, err := utils.GetRepoSigInfo(cveIssueRaw.Repository.Name)
	if err != nil {
		return []string{}
	}
	sigMaintainers, _, err := utils.GetMembersBySig(sigGroup)
	if err != nil {
		return []string{}
	}

	repo := strings.Split(cveIssueRaw.Repository.FullName, "/")
	repoAdmins, err := utils.GetAllAdmins(repo[0], repo[1])
	if err != nil {
		return []string{}
	}
	followUsers := slices.Concat(sigMaintainers, repoAdmins)
	followUsers = utils.Difference(followUsers, cveIssueRaw.todoUsers())
	followUsers = utils.Difference(followUsers, cveIssueRaw.applyUsers())
	return followUsers
}

func (cveIssueRaw *CVEIssueRaw) GetFollowUsers(events CloudEvents) {
	events.SetExtension("followusers", strings.Join(cveIssueRaw.followUsers(), ","))
}

func (cveIssueRaw *CVEIssueRaw) todoUsers() []string {
	var todoUsers []string
	if cveIssueRaw.Issue.Assignee != nil {
		todoUsers = []string{cveIssueRaw.Issue.Assignee.UserName}
	} else {
		todoUsers = []string{}
	}
	todoUsers = utils.Difference(todoUsers, cveIssueRaw.applyUsers())
	return todoUsers
}

func (cveIssueRaw *CVEIssueRaw) GetTodoUsers(events CloudEvents) {
	events.SetExtension("todousers", strings.Join(cveIssueRaw.todoUsers(), ","))
	events.SetExtension("businessid", strconv.Itoa(int(cveIssueRaw.Issue.Id)))
}

func (cveIssueRaw *CVEIssueRaw) applyUsers() []string {
	return []string{}
}

func (cveIssueRaw *CVEIssueRaw) GetApplyUsers(events CloudEvents) {
	events.SetExtension("applyusers", strings.Join(cveIssueRaw.applyUsers(), ","))
}

func (cveIssueRaw *CVEIssueRaw) IsDone(events CloudEvents) {
	events.SetExtension("isdone", slices.Contains([]string{"rejected", "closed"},
		cveIssueRaw.Issue.State))
}

// ToMap transfer cve issue raw to map[string]interface{}.
func (cveIssueRaw *CVEIssueRaw) ToMap() RawMap {
	cveMap := extractVariables(*cveIssueRaw.Description)
	cveIssueMap := StructToMap(cveIssueRaw)
	for s, i := range cveMap {
		cveIssueMap[s] = i
	}
	return cveIssueMap
}

func extractVariables(text string) map[string]interface{} {
	result := make(map[string]interface{})

	// 定义正则表达式来匹配每个变量
	patterns := map[string]string{
		"CVENumber":        `漏洞编号：(.*?)\n`,
		"CVEComponent":     `漏洞归属组件：\[(.*?)]`,
		"CVRVersion":       `漏洞归属的版本：((?s).*?)\n`,
		"CVEBaseScore":     `BaseScore：(.*?)\n`,
		"CVEVector":        `Vector：(.*?)\n`,
		"CVEDesc":          `漏洞简述：(.*?)漏洞公开时间`,
		"CVEReleaseDate":   `漏洞公开时间：(.*?)\n`,
		"CVECreatedDate":   `漏洞创建时间：(.*?)\n`,
		"CVEDetailURL":     `漏洞详情参考链接：(.*?)\n`,
		"CVEAffectVersion": `受影响版本排查\(受影响/不受影响\)：((?s).*?)\n\n`,
		"CVEApiChange":     `修复是否涉及abi变化\(是/否\)：((?s).*?)\n\n`,
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
