package i18n

import (
	"encoding/json"
	"fmt"

	i18nLib "github.com/nicksnyder/go-i18n/v2/i18n"
	"github.com/scienceandcode/habits-todo-backend/pkg/common"
	"golang.org/x/text/language"
)

var localizer *i18nLib.Localizer
var bundle *i18nLib.Bundle

func Init() {
	bundle = i18nLib.NewBundle(language.English) // Default language
	bundle.RegisterUnmarshalFunc("json", json.Unmarshal)

	bundle.MustLoadMessageFile(fmt.Sprintf("../../internal/i18n/resources/%s.json", common.GetEnv("I18N_LANGUAGE")))

	localizer = i18nLib.NewLocalizer(bundle, common.GetEnv("I18N_LANGUAGE"))
}

func Message(key string) string {
	return localizer.MustLocalize(&i18nLib.LocalizeConfig{
		MessageID: key,
	})
}
