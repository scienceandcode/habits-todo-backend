package i18n

import (
	"github.com/scienceandcode/habits-todo-backend/internal/i18n/messages"
)

func getSupportedLanguages() map[string]map[string]string {
	msg := map[string]map[string]string{}

	msg["pt_br"] = messages.PtBr()
	msg["en_us"] = messages.EnUs()

	return msg
}
