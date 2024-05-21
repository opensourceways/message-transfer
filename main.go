package main

import (
	"github.com/opensourceways/server-common-lib/logrusutil"
	"github.com/sirupsen/logrus"
	"message-transfer/common/kafka"
	"message-transfer/common/postgresql"
	"message-transfer/config"
	"message-transfer/service"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt, syscall.SIGTERM)

	logrusutil.ComponentInit("messageAdapter-collect")
	log := logrus.NewEntry(logrus.StandardLogger())

	cfg := new(config.Config)
	initConfig(cfg)

	defer kafka.Exit()
	if err := postgresql.Init(&cfg.Postgresql, false); err != nil {
		logrus.Errorf("init postgresql failed, err:%s", err.Error())
		return
	}

	if err := kafka.Init(&cfg.Kafka, log, false); err != nil {
		logrus.Errorf("init kafka failed, err:%s", err.Error())
		return
	}

	go func() {
		service.SubscribeEurRaw()
	}()
	<-sig
}

func initConfig(cfg *config.Config) {
	pgCfg := postgresql.NewTestConfig()
	pgCfg.SetDefault()
	cfg.Postgresql = pgCfg
	cfg.Kafka.SetDefault()
}
