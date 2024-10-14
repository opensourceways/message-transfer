/*
Copyright (c) Huawei Technologies Co., Ltd. 2024. All rights reserved
*/

// Package service handle func.
package service

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/sirupsen/logrus"

	"github.com/opensourceways/message-transfer/common/kafka"
	"github.com/opensourceways/message-transfer/config"
	"github.com/opensourceways/message-transfer/models/dto"
	"github.com/opensourceways/message-transfer/utils"
)

func handle(raw dto.Raw, cfg kafka.ConsumeConfig) error {
	time.Sleep(utils.GetConsumeSleepTime())
	event := raw.ToCloudEventByConfig(cfg.Topic)
	if event.ID() == "" {
		return nil
	}
	err := event.SaveDb()
	if err != nil {
		logrus.Errorf("saveDb failed, err:%v", err)
		return err
	}
	kafkaSendErr := kafka.SendMsg(cfg.Publish, &event)
	if kafkaSendErr != nil {
		return kafkaSendErr
	}
	return nil
}

// CVEHandle handle cve issue raw.
func CVEHandle(payload []byte, _ map[string]string) error {
	var raw dto.CVEIssueRaw
	msgBodyErr := json.Unmarshal(payload, &raw)
	if msgBodyErr != nil {
		return msgBodyErr
	}
	sigGroupName, err := utils.GetRepoSigInfo(raw.Repository.Name)
	if err != nil {
		return err
	}
	sigMaintainers, _, err := utils.GetMembersBySig(sigGroupName)
	if err != nil {
		return err
	}

	raw.SigGroupName = sigGroupName
	raw.SigMaintainers = sigMaintainers
	rawMap := raw.ToMap()
	return handle(rawMap, config.CveConfigInstance.Kafka)
}

// GiteeIssueHandle handle gitee issue raw.
func GiteeIssueHandle(payload []byte, _ map[string]string) error {
	var raw dto.GiteeIssueRaw
	msgBodyErr := json.Unmarshal(payload, &raw)
	if msgBodyErr != nil {
		return msgBodyErr
	}
	sigGroupName, err := utils.GetRepoSigInfo(raw.Repository.Name)
	if err != nil {
		return err
	}
	sigMaintainers, _, err := utils.GetMembersBySig(sigGroupName)
	if err != nil {
		return err
	}
	repo := strings.Split(raw.Repository.FullName, "/")
	repoAdmins, err := utils.GetAllAdmins(repo[0], repo[1])
	if err != nil {
		return err
	}

	raw.SigGroupName = sigGroupName
	raw.SigMaintainers = sigMaintainers
	raw.RepoAdmins = repoAdmins
	rawMap := dto.StructToMap(raw)
	return handle(rawMap, config.GiteeConfigInstance.Issue)
}

// GiteePushHandle handle gitee push raw.
func GiteePushHandle(payload []byte, _ map[string]string) error {
	var raw dto.GiteePushRaw
	msgBodyErr := json.Unmarshal(payload, &raw)
	if msgBodyErr != nil {
		return msgBodyErr
	}
	sigGroupName, err := utils.GetRepoSigInfo(raw.Repository.Name)
	if err != nil {
		return err
	}

	sigMaintainers, _, err := utils.GetMembersBySig(sigGroupName)
	if err != nil {
		return err
	}
	repo := strings.Split(raw.Repository.FullName, "/")
	repoAdmins, err := utils.GetAllAdmins(repo[0], repo[1])
	if err != nil {
		return err
	}

	raw.SigGroupName = sigGroupName
	raw.SigMaintainers = sigMaintainers
	raw.RepoAdmins = repoAdmins
	rawMap := dto.StructToMap(raw)
	return handle(rawMap, config.GiteeConfigInstance.Push)
}

// GiteePrHandle handle gitee pr raw.
func GiteePrHandle(payload []byte, _ map[string]string) error {
	var raw dto.GiteePrRaw
	msgBodyErr := json.Unmarshal(payload, &raw)
	if msgBodyErr != nil {
		return msgBodyErr
	}
	sigGroupName, err := utils.GetRepoSigInfo(raw.Repository.Name)
	if err != nil {
		return err
	}
	sigMaintainers, _, err := utils.GetMembersBySig(sigGroupName)
	if err != nil {
		return err
	}
	repo := strings.Split(raw.Repository.FullName, "/")
	repoAdmins, err := utils.GetAllAdmins(repo[0], repo[1])
	if err != nil {
		return err
	}

	raw.SigGroupName = sigGroupName
	raw.SigMaintainers = sigMaintainers
	raw.RepoAdmins = repoAdmins
	rawMap := dto.StructToMap(raw)
	return handle(rawMap, config.GiteeConfigInstance.PR)
}

// GiteeNoteHandle handle gitee note raw.
func GiteeNoteHandle(payload []byte, _ map[string]string) error {
	var raw dto.GiteeNoteRaw
	msgBodyErr := json.Unmarshal(payload, &raw)
	if msgBodyErr != nil {
		return msgBodyErr
	}
	sigGroupName, err := utils.GetRepoSigInfo(raw.Repository.Name)
	if err != nil {
		return err
	}
	sigMaintainers, _, err := utils.GetMembersBySig(sigGroupName)
	if err != nil {
		return err
	}

	repo := strings.Split(raw.Repository.FullName, "/")
	repoAdmins, err := utils.GetAllAdmins(repo[0], repo[1])
	if err != nil {
		return err
	}

	raw.SigGroupName = sigGroupName
	raw.SigMaintainers = sigMaintainers
	raw.RepoAdmins = repoAdmins
	rawMap := dto.StructToMap(raw)
	return handle(rawMap, config.GiteeConfigInstance.Note)
}

// EurBuildHandle handle eur build raw.
func EurBuildHandle(payload []byte, _ map[string]string) error {
	var raw dto.EurBuildMessageRaw
	msgBodyErr := json.Unmarshal(payload, &raw)
	if msgBodyErr != nil {
		return msgBodyErr
	}
	raw.SourceGroup = fmt.Sprintf("%s/%s", raw.Body.Owner, raw.Body.Copr)
	rawMap := dto.StructToMap(raw)
	return handle(rawMap, config.EurBuildConfigInstance.Kafka)
}

// OpenEulerMeetingHandle handle openEuler meeting raw.
func OpenEulerMeetingHandle(payload []byte, _ map[string]string) error {
	var raw dto.OpenEulerMeetingRaw
	msgBodyErr := json.Unmarshal(payload, &raw)
	if msgBodyErr != nil {
		logrus.Errorf("unmarshal meeting message failed, err:%v", msgBodyErr)
		return msgBodyErr
	}

	raw.MeetingStartTime = raw.Msg.Date + raw.Msg.Start
	raw.MeetingEndTime = raw.Msg.Date + raw.Msg.End
	raw.Time = time.Now()
	sigMaintainers, _, err := utils.GetMembersBySig(raw.Msg.GroupName)
	if err != nil {
		return err
	}
	raw.SigMaintainers = sigMaintainers
	rawMap := dto.StructToMap(raw)
	return handle(rawMap, config.MeetingConfigInstance.Kafka)
}

func ForumHandle(payload []byte, _ map[string]string) error {
	var raw dto.Notification
	msgBodyErr := json.Unmarshal(payload, &raw)
	if msgBodyErr != nil {
		return msgBodyErr
	}
	rawMap := dto.StructToMap(raw)
	return handle(rawMap, config.ForumConfigInstance.Kafka)
}
