package logger

import (
	"net"
	"strings"

	"github.com/bshuster-repo/logrus-logstash-hook"
	log "github.com/sirupsen/logrus"
)

func initLogstashlogger(config LogConfig) *log.Logger {
	logger := log.New()
	conn, err := net.Dial(config.NetworkType, config.Host+":"+config.Port)
	if err != nil {
		log.Fatal(err)
	}
	hook := logrustash.New(conn, logrustash.DefaultFormatter(log.Fields{"type": strings.ToLower(config.Title)}))
	if err != nil {
		log.Fatal(err)
	}
	logger.Hooks.Add(hook)
	//logger.Out = ioutil.Discard
	return logger
}
