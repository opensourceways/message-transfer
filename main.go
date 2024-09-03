package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/opensourceways/message-transfer/common/kafka"
	"github.com/opensourceways/message-transfer/common/postgresql"
	"github.com/opensourceways/message-transfer/config"
	"github.com/opensourceways/message-transfer/service"
	"github.com/opensourceways/message-transfer/utils"
	"github.com/opensourceways/server-common-lib/logrusutil"
	"github.com/sirupsen/logrus"
)

func main() {
	logrusutil.ComponentInit("message-transfer")
	log := logrus.NewEntry(logrus.StandardLogger())

	cfg := initConfig()

	logrus.Info("初始化pg,配置文件:", cfg.Postgresql)
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

	//go func() {
	//	service.SubscribeEurRaw()
	//}()
	//go func() {
	//	service.SubscribeGiteeIssue()
	//}()
	//go func() {
	//	service.SubscribeGiteePush()
	//}()
	//go func() {
	//	service.SubscribeGiteePr()
	//}()
	//go func() {
	//	service.SubscribeGiteeNote()
	//}()
	//go func() {
	//	service.SubscribeOpenEulerMeeting()
	//}()
	go func() {
		service.SubscribeCVERaw()
	}()

	select {}
}

func initConfig() *config.Config {
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
	return cfg
}

func initTransferConfig(o Options) {
	config.InitGiteeConfig(o.GiteeConfig)
	config.InitEurBuildConfig(o.EurBuildConfig)
	config.InitMeetingConfig(o.OpenEulerMeetingConfig)
	config.InitCVEConfig(o.CVEConfig)
}

/*
获取启动参数，配置文件地址由启动参数传入
*/
func gatherOptions(fs *flag.FlagSet, args ...string) (Options, error) {
	var o Options
	fmt.Println("从环境变量接收参数", args)
	o.AddFlags(fs)
	err := fs.Parse(args)
	return o, err
}

type Options struct {
	Config                 string
	EurBuildConfig         string
	GiteeConfig            string
	OpenEulerMeetingConfig string
	CVEConfig              string
}

func (o *Options) AddFlags(fs *flag.FlagSet) {
	fs.StringVar(&o.Config, "config-file", "", "Path to config file.")
	fs.StringVar(&o.EurBuildConfig, "eur-build-config-file", "", "Path to eur-build config file.")
	fs.StringVar(&o.GiteeConfig, "gitee-config-file", "", "Path to gitee config file.")
	fs.StringVar(&o.OpenEulerMeetingConfig, "meeting-config-file", "", "Path to gitee config file.")
	fs.StringVar(&o.CVEConfig, "cve-config-file", "", "Path to gitee config file.")

}
