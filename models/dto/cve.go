/*
Copyright (c) Huawei Technologies Co., Ltd. 2024. All rights reserved
*/

// Package dto models dto of cve
package dto

import (
	"regexp"
	"slices"
	"strings"

	"github.com/opensourceways/go-gitee/gitee"

	"github.com/opensourceways/message-transfer/utils"
)

// CVEIssueRaw cve issue raw.
type CVEIssueRaw struct {
	gitee.IssueEvent
}

func (cveIssueRaw *CVEIssueRaw) ToCloudEventsByConfig() CloudEvents {
	rawMap := cveIssueRaw.ToMap()
	return rawMap.ToCloudEventByConfig("cve_issue_raw")
}

func (cveIssueRaw *CVEIssueRaw) GetRelateUsers(events CloudEvents) {
	events.SetExtension("relatedusers", []string{})
}

func (cveIssueRaw *CVEIssueRaw) GetFollowUsers(events CloudEvents) {
	sigGroup, err := utils.GetRepoSigInfo(cveIssueRaw.Repository.Name)
	if err != nil {
		return
	}
	sigMaintainers, _, err := utils.GetMembersBySig(sigGroup)
	if err != nil {
		return
	}

	repo := strings.Split(cveIssueRaw.Repository.FullName, "/")
	repoAdmins, err := utils.GetAllAdmins(repo[0], repo[1])
	if err != nil {
		return
	}
	followUsers := slices.Concat(sigMaintainers, repoAdmins)
	var assignee []string
	if cveIssueRaw.Issue.Assignee != nil {
		assignee = []string{cveIssueRaw.Issue.Assignee.UserName}
	}
	followUsers = utils.Difference(followUsers, assignee)
	events.SetExtension("followusers", followUsers)
}

func (cveIssueRaw *CVEIssueRaw) GetTodoUsers(events CloudEvents) {
	var todoUsers []string
	if cveIssueRaw.Issue.Assignee != nil {
		todoUsers = []string{cveIssueRaw.Issue.Assignee.UserName}
	} else {
		todoUsers = []string{}
	}
	events.SetExtension("todousers", todoUsers)
	events.SetExtension("businessid", cveIssueRaw.Issue.Id)
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
