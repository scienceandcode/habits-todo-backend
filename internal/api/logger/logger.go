package logger

import (
	"log"

	"github.com/scienceandcode/habits-todo-backend/pkg/common"
)

func Info(message string) {
	log.Println(getLogMessage("INFO", message))
}

func Error(message string) {
	log.Println(getLogMessage("ERROR", message))
}

func Fatal(message string) {
	log.Fatalln(getLogMessage("FATAL", message))
}

func Panic(message string) {
	log.Panicln(getLogMessage("PANIC", message))
}

func Debug(message string) {
	if common.GetEnv("GIN_MODE") == "debug" {
		log.Println(getLogMessage("DEBUG", message))
	}
}

func Warn(message string) {
	log.Println(getLogMessage("WARN", message))
}

func getLogMessage(logType, message string) string {
	return "[" + common.GetEnv("APPLICATION_NAME") + "]" + "[" + logType + "] " + message
}
