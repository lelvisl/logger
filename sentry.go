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

func sentryHook(config *LogConfig) (*logrus_sentry.SentryHook, error) {
	var (
		hook *logrus_sentry.SentryHook
		err  error
	)
	levels := logLevels(config)
	client, err := raven.New(config.Sentry.DSN)
	if err != nil {
		return nil, err
	}
	if len(config.Sentry.Tags) != 0 {
		client.Tags = config.Sentry.Tags
	}

	hook, err = logrus_sentry.NewWithClientSentryHook(client, levels)

	return hook, err
}
