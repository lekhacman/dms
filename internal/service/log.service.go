package service

import log "github.com/sirupsen/logrus"

func NewLogger(level string) *log.Logger {
	logger := log.New()
	logMap := map[string]log.Level{
		"panic": log.PanicLevel,
		"fatal": log.FatalLevel,
		"error": log.ErrorLevel,
		"warn":  log.WarnLevel,
		"info":  log.InfoLevel,
		"debug": log.DebugLevel,
		"trace": log.TraceLevel,
	}
	logLevel, ok := logMap[level]
	if !ok {
		log.Fatal("Log level not found")
	}
	logger.SetLevel(logLevel)

	return logger
}
