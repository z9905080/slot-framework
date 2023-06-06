package logger

import (
	"github.com/z9905080/gloger"
	"slot-framework/environment"
	"strings"
)

func NewLogger(config environment.Config) Logger {

	log := gloger.NewLogger()

	// create log level mapping
	logLevelMapping := map[string]gloger.Level{
		"debug":   gloger.DEBUG,
		"info":    gloger.INFO,
		"warning": gloger.WARNING,
		"error":   gloger.ERROR,
		"fatal":   gloger.FATAL,
	}

	log.SetCurrentLevel(logLevelMapping[strings.ToLower(config.LogSetting.Level)])

	// create log mode mapping
	logModeMapping := map[string]gloger.OutputMode{
		"stdout": gloger.Stdout,
		"file":   gloger.File,
	}

	log.SetLogMode(logModeMapping[strings.ToLower(config.LogSetting.Output)])

	// set caller depth
	//log.SetCallerDepth(config.LogSetting.Depth)
	return log
}
