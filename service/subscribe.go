/*
Copyright (c) Huawei Technologies Co., Ltd. 2024. All rights reserved
*/

package service

import (
	kfklib "github.com/opensourceways/kafka-lib/agent"
	"github.com/opensourceways/message-transfer/config"
	"github.com/sirupsen/logrus"
)

// SubscribeEurRaw subscribe eur topic.
func SubscribeEurRaw() {
	logrus.Info("subscribing to eur topic")
	_ = kfklib.Subscribe(config.EurBuildConfigInstance.Kafka.Group, EurBuildHandle,
		[]string{config.EurBuildConfigInstance.Kafka.Topic})
}

// SubscribeGiteeIssue subscribe gitee issue topic.
func SubscribeGiteeIssue() {
	logrus.Info("subscribing to issue topic")
	_ = kfklib.Subscribe(config.GiteeConfigInstance.Issue.Group, GiteeIssueHandle,
		[]string{config.GiteeConfigInstance.Issue.Topic})
}

// SubscribeGiteePr subscribe gitee pullRequest topic.
func SubscribeGiteePr() {
	logrus.Info("subscribing to pr topic")
	_ = kfklib.Subscribe(config.GiteeConfigInstance.PR.Group, GiteePrHandle,
		[]string{config.GiteeConfigInstance.PR.Topic})
}

// SubscribeGiteeNote subscribe gitee note topic.
func SubscribeGiteeNote() {
	logrus.Info("subscribing to note topic")
	_ = kfklib.Subscribe(config.GiteeConfigInstance.Note.Group, GiteeNoteHandle,
		[]string{config.GiteeConfigInstance.Note.Topic})
}

// SubscribeOpenEulerMeeting subscribe meeting topic.
func SubscribeOpenEulerMeeting() {
	logrus.Info("subscribing to openEuler meeting topic")
	_ = kfklib.Subscribe(config.MeetingConfigInstance.Kafka.Group, OpenEulerMeetingHandle,
		[]string{config.MeetingConfigInstance.Kafka.Topic})
}

// SubscribeCVERaw subscribe cve topic.
func SubscribeCVERaw() {
	logrus.Info("subscribing to cve topic")
	_ = kfklib.Subscribe(config.CveConfigInstance.Kafka.Group, CVEHandle,
		[]string{config.CveConfigInstance.Kafka.Topic})
}

// SubscribeForumRaw subscribe forum topic.
func SubscribeForumRaw() {
	logrus.Info("subscribing to openEuler forum topic")
	_ = kfklib.Subscribe(config.ForumConfigInstance.Kafka.Group, OpenEulerForumHandle,
		[]string{config.ForumConfigInstance.Kafka.Topic})
}
