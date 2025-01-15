package logger

import (
	log "github.com/sirupsen/logrus"
)

func setupFormatter() {
	log.SetFormatter(&log.JSONFormatter{})
}

func Info(message string) {
	setupFormatter()
	log.Info(message)
}

func Error(message string) {
	setupFormatter()
	log.Error(message)
}

func Debug(message string) {
	setupFormatter()
	log.Debug(message)
}

func Warn(message string) {
	setupFormatter()
	log.Warn(message)
}
