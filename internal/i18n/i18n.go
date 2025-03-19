package i18n

import (
	i18nLib "github.com/nicksnyder/go-i18n/v2/i18n"
	"github.com/scienceandcode/habits-todo-backend/pkg/common"
	"golang.org/x/text/language"
)

var localizer *i18nLib.Localizer
var bundle *i18nLib.Bundle

func Init() {
	translationMessages := getSupportedLanguages()
	envLanguage := common.GetEnv("I18N_LANGUAGE")

	messages, ok := translationMessages[envLanguage]

	if !ok {
		panic("Variable I18N_LANGUAGE=" + envLanguage + " not supported. Please check the supported languages in internal/i18n/supported_languages.go")
	}

	languageTagString := messages["language"]
	languageTag, _ := language.Parse(languageTagString)

	bundle = i18nLib.NewBundle(languageTag)

	for key, message := range messages {
		err := bundle.AddMessages(languageTag, &i18nLib.Message{ID: key, Other: message})
		if err != nil {
			panic(err)
		}
	}

	localizer = i18nLib.NewLocalizer(bundle, common.GetEnv("I18N_LANGUAGE"))
}

func Message(key string) string {
	return localizer.MustLocalize(&i18nLib.LocalizeConfig{
		MessageID: key,
	})
}
