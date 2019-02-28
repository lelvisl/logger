package logger

import (
	"context"
	"github.com/onrik/logrus/filename"
	log "github.com/sirupsen/logrus"
	"os"
)

//LogConfig is a configuration for logger
type LogConfig struct {
	Title       string       `yaml:"title" json:"title" toml:"title"`
	Type        string       `yaml:"type" json:"type" toml:"type"`
	NetworkType string       `yaml:"network type" json:"network_type" toml:"network_type"`
	Host        string       `yaml:"host" json:"host" toml:"host"`
	Severity    string       `yaml:"severity" json:"severity" toml:"severity"`
	Facility    string       `yaml:"facility" json:"facility" toml:"facility"`
	Port        string       `yaml:"port" json:"port" toml:"port"`
	FilePath    string       `yaml:"file path" json:"file_path" toml:"file_path"`
	FileName    string       `yaml:"file name" json:"file_name" toml:"file_name"`
	DebugMode   bool         `yaml:"debug mode" json:"debug_mode" toml:"debug_mode"`
	Sentry      sentryConfig `yaml: "sentry" json:"sentry"`
}

type sentryConfig struct {
	Tags map[string]string `yaml:"tags" json:"tags"`
	DSN  string            `yaml:"dsn" json:"dns"`
}

type ctxlog struct{}

//WithLogger put logger to context
func WithLogger(ctx context.Context, logger *log.Logger) context.Context {
	return context.WithValue(ctx, ctxlog{}, logger)
}

//WithEntry put logger.Entry to context
func WithEntry(ctx context.Context, logger *log.Entry) context.Context {
	return context.WithValue(ctx, ctxlog{}, logger)
}

//Logger get logger from context
func Logger(ctx context.Context) *log.Logger {
	//	l, _ := ctx.Value("logger").(*log.Logger)
	l, ok := ctx.Value(ctxlog{}).(*log.Logger)
	if !ok {
		//return DefaultLogger
		l = initLogger(LogConfig{Type: "stdout", Severity: "LOG_INFO"})
	}
	return l
}

//Entry get logger from context
func Entry(ctx context.Context) *log.Entry {
	//	l, _ := ctx.Value("logger").(*log.Entry)
	l, ok := ctx.Value(ctxlog{}).(*log.Entry)
	if !ok {
		//return DefaultLogger
		l = initLogger(LogConfig{Type: "stdout", Severity: "LOG_INFO"}).WithField("", "")
	}
	return l
}

//CreateLogger from config
func CreateLogger(config LogConfig) *log.Logger {
	logger := initLogger(config)
	return logger
}

func initLogger(config LogConfig) *log.Logger {
	logger := log.New()
	filenameHook := filename.NewHook()
	filenameHook.Field = "source" // Customize source field name
	logger.AddHook(filenameHook)
	switch config.Type {
	case "syslog":
		logger = initSyslogger(config)
	case "logstash":
		logger = initLogstashlogger(config)
	case "stdout":
		logger.Out = os.Stdout
		return logger
	case "stderr":
		logger.Out = os.Stderr
		return logger
	case "sentry":
		logger = initSentrylogger(config)
	default:
	}
	if config.DebugMode {
		logger.Out = os.Stdout
	}
	logger.Formatter = &log.TextFormatter{}
	logger.Level = logLevel[config.Severity]
	return logger
}

var logLevel = map[string]log.Level{
	"LOG_EMERG":   log.PanicLevel,
	"LOG_ALERT":   log.PanicLevel,
	"LOG_CRIT":    log.FatalLevel,
	"LOG_ERR":     log.ErrorLevel,
	"LOG_WARNING": log.WarnLevel,
	"LOG_NOTICE":  log.InfoLevel,
	"LOG_INFO":    log.InfoLevel,
	"LOG_DEBUG":   log.DebugLevel,
}
