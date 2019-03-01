package logger

import (
	"testing"
)

//
func TestCreateLogger(t *testing.T) {
	type testCase struct {
		Message  string
		LogLevel string
	}

	testCases := []testCase{
		{
			"Info text",
			"LOG_INFO",
		},
		{
			"Warn text",
			"LOG_WARNING",
		},
		{
			"Debug text",
			"LOG_DEBUG",
		},
	}

	for _, tc := range testCases {
		c := LogConfig{
			Type:     "stdout",
			Severity: tc.LogLevel,
			Sentry: SentryConfig{
				Tags: map[string]string{
					"site": "dev",
				},
				// TODO insert DNS for raven
				DSN: "",
			},
		}
		logger := CreateLogger(c)
		logger.Infoln(tc.LogLevel, "Info text")
		logger.Warningln(tc.LogLevel, " Warn text")
		logger.Debugln(tc.LogLevel, "Debug text")
	}
}
