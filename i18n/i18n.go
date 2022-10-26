package i18n

import (
	"log"
	"os"

	"github.com/nicksnyder/go-i18n/v2/i18n"
	"go.uber.org/fx"
	"golang.org/x/text/language"
)

var bundle *i18n.Bundle

func NewI18n(localesDirs ...string) {
	localesDir := "./locales"
	if len(localesDirs) > 0 {
		localesDir = localesDirs[0]
	}
	files, err := os.ReadDir(localesDir)
	if err != nil {
		log.Fatal(err)
	}

	bundle = i18n.NewBundle(language.English)
	for _, file := range files {
		filename := file.Name()
		data, err := os.ReadFile(localesDir + "/" + filename)
		if err != nil {
			log.Fatal(err)
		}
		bundle.MustParseMessageFileBytes(data, filename)
	}
}

var Module = fx.Invoke(NewI18n)

func Translate(key string, tags ...language.Tag) string {
	tag := language.English
	if len(tags) > 0 {
		tag = tags[0]
	}
	value, err := i18n.NewLocalizer(bundle, tag.String()).Localize(
		&i18n.LocalizeConfig{
			MessageID: key,
		},
	)
	if err != nil {
		return key
	}
	return value
}
