package factory

import "translator_factory/models"

const (
	German = iota + 1
	English
	Japanese
)

type Translator interface {
	Translate() string
}

func CreateTranslator(languateType int) Translator {
	switch languateType {
	case German:
		return new(models.GermanTranslator)
	case English:
		return new(models.EnglishTranslator)
	case Japanese:
		return new(models.JapaneseTranslator)
	default:
		panic("no such translator")
	}
}
