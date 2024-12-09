/*
Copyright (c) Huawei Technologies Co., Ltd. 2024. All rights reserved
*/

package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/opensourceways/server-common-lib/logrusutil"
	"github.com/sirupsen/logrus"

	"github.com/opensourceways/message-transfer/common/kafka"
	"github.com/opensourceways/message-transfer/common/postgresql"
	"github.com/opensourceways/message-transfer/config"
	"github.com/opensourceways/message-transfer/service"
	"github.com/opensourceways/message-transfer/utils"
)

func main() {
	logrus.SetFormatter(&logrus.JSONFormatter{
		PrettyPrint: true, // 使 JSON 输出更美观
	})
	logrusutil.ComponentInit("message-transfer")
	log := logrus.NewEntry(logrus.StandardLogger())

	cfg, projectName := initConfig()

	if err := postgresql.Init(&cfg.Postgresql, false); err != nil {
		logrus.Errorf("init postgresql failed, err:%s", err.Error())
		return
	}

	if err := kafka.Init(&cfg.Kafka, log, false); err != nil {
		logrus.Errorf("init kafka failed, err:%s", err.Error())
		return
	}

	if err := utils.Init(&cfg.Utils); err != nil {
		logrus.Errorf("init utils failed, err:%s", err.Error())
		return
	}

	switch projectName {
	case "openEuler":
		subscribeOpenEuler()
	case "openUBMC":
		subscribeOpenUBMC()
	}
}

func subscribeOpenEuler() {
	go func() {
		service.SubscribeEurRaw()
	}()
	go func() {
		service.SubscribeGiteeIssue()
	}()
	go func() {
		service.SubscribeGiteePr()
	}()
	go func() {
		service.SubscribeGiteeNote()
	}()
	go func() {
		service.SubscribeOpenEulerMeeting()
	}()
	go func() {
		service.SubscribeCVERaw()
	}()
	go func() {
		service.SubscribeForumRaw()
	}()

	select {}
}

func subscribeOpenUBMC() {

}

func initConfig() (*config.Config, string) {
	o, err := gatherOptions(
		flag.NewFlagSet(os.Args[0], flag.ExitOnError),
		os.Args[1:]...,
	)
	if err != nil {
		logrus.Fatalf("new Options failed, err:%s", err.Error())
	}
	cfg := new(config.Config)

	if err := utils.LoadFromYaml(o.Config, cfg); err != nil {
		logrus.Error("Config初始化失败, err:", err)
	}
	initTransferConfig(o)
	return cfg, o.ProjectName
}

func initTransferConfig(o options) {
	switch o.ProjectName {
	case "openEuler":
		initOpenEulerTransferConfig(o)
	case "openUBMC":
		initOpenUBMCTransferConfig(o)
	}
}

func initOpenEulerTransferConfig(o options) {
	config.InitGiteeConfig(o.GiteeConfig)
	config.InitEurBuildConfig(o.EurBuildConfig)
	config.InitMeetingConfig(o.MeetingConfig)
	config.InitCVEConfig(o.CVEConfig)
	config.InitForumConfig(o.ForumConfig)
}

func initOpenUBMCTransferConfig(o options) {
	config.InitMeetingConfig(o.MeetingConfig)
}

/*
获取启动参数，配置文件地址由启动参数传入
*/
func gatherOptions(fs *flag.FlagSet, args ...string) (options, error) {
	var o options
	fmt.Println("从环境变量接收参数", args)
	o.AddFlags(fs)
	err := fs.Parse(args)
	return o, err
}

type options struct {
	Config         string
	EurBuildConfig string
	GiteeConfig    string
	MeetingConfig  string
	CVEConfig      string
	ForumConfig    string
	ProjectName    string
}

// AddFlags add flags.
func (o *options) AddFlags(fs *flag.FlagSet) {
	fs.StringVar(&o.Config, "config-file", "", "Path to config file.")
	fs.StringVar(&o.EurBuildConfig, "eur-build-config-file", "", "Path to eur-build config file.")
	fs.StringVar(&o.GiteeConfig, "gitee-config-file", "", "Path to gitee config file.")
	fs.StringVar(&o.MeetingConfig, "meeting-config-file", "",
		"Path to meeting config file.")
	fs.StringVar(&o.CVEConfig, "cve-config-file", "", "Path to cve config file.")
	fs.StringVar(&o.ForumConfig, "forum-config-file", "", "Path to forum config file.")
	fs.StringVar(&o.ProjectName, "project-name", "", "Project name")
}
