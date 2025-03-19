package messages

import "golang.org/x/text/language"

func EnUs() map[string]string {
	return map[string]string{
		"language": language.AmericanEnglish.String(),

		"Admin.MigrateModels.Response.Success": "Models migrated successfully",
	}
}
