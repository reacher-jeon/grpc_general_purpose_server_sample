package logger_test

import (
	"fmt"
	"task-root/internal/util/logger"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLogLevelName(t *testing.T) {
	names := []string{
		"DISABLE",
		"EMERGENCY",
		"ALERT",
		"CRITICAL",
		"ERROR",
		"WARNING",
		"NOTICE",
		"INFO",
		"DEBUG",
		"DEFAULT",
		"TRACE",
	}

	for level := logger.LogLevelDisable; level <= logger.LogLevelTrace; level++ {
		res := logger.LogLevelToString[level]
		fmt.Printf("OK name : [%v] : Return name : [%v]\n", names[level], res)
		assert.Equal(t, names[level], res)
	}
}

func TestLogLeveSetFromString(t *testing.T) {
	names := []string{
		"DISABLE",
		"EMERGENCY",
		"ALERT",
		"CRITICAL",
		"ERROR",
		"WARNING",
		"NOTICE",
		"INFO",
		"DEBUG",
		"DEFAULT",
		"TRACE",
	}

	for level := logger.LogLevelDisable; level <= logger.LogLevelTrace; level++ {
		res := logger.GetLogLevelFromString(names[level])
		fmt.Printf("OK name : [%v[%v]], Return index : [%v]\n", names[level], level, res)
		assert.Equal(t, logger.LogLevel(level), res)
	}
}

func TestNewLogger(t *testing.T) {
	log := logger.New(logger.LogLevelTrace)
	log.Trace().Msg("testLogLevel Trace")
}
