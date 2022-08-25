package i18n

import (
	"log"
	"os"

	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
)

const LOCALES_DIR = "./locales"

var bundle *i18n.Bundle

func init() {
	files, err := os.ReadDir(LOCALES_DIR)
	if err != nil {
		log.Fatal(err)
	}

	bundle = i18n.NewBundle(language.English)
	for _, file := range files {
		filename := file.Name()
		data, err := os.ReadFile(LOCALES_DIR + "/" + filename)
		if err != nil {
			log.Fatal(err)
		}
		bundle.MustParseMessageFileBytes(data, filename)
	}
}

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
