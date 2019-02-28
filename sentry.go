package logger

import (
	"github.com/evalphobia/logrus_sentry"
	"github.com/getsentry/raven-go"
	log "github.com/sirupsen/logrus"
)

func logLevels(config *LogConfig) []log.Level {
	// TODO must
	ll := logLevel[config.Severity]
	levels := make([]log.Level, 0, ll+1)
	for i := log.PanicLevel; i <= ll; i++ {
		levels = append(levels, i)
	}
	return levels
}

func initSentrylogger(config LogConfig) *log.Logger {
	var (
		err error
	)
	logger := log.New()
	_, err = raven.New(config.Sentry.DSN)
	if err != nil {
		log.Fatal(err)
	}

	hook, err := sentryHook(&config)
	if err == nil {
		logger.Hooks.Add(hook)
	}

	return logger
}

func sentryHook(config *LogConfig) (*logrus_sentry.SentryHook, error) {
	var (
		hook *logrus_sentry.SentryHook
		err  error
	)
	if len(config.Sentry.Tags) != 0 {
		hook, err = logrus_sentry.NewWithTagsSentryHook(config.Sentry.DSN, config.Sentry.Tags, logLevels(config))
	}

	hook, err = logrus_sentry.NewSentryHook(config.Sentry.DSN, logLevels(config))

	return hook, err
}
