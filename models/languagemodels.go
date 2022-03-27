package models

//德语翻译器
type GermanTranslator struct{}

func (*GermanTranslator) Translate() string {
	return "这是德语翻译器"
}

//英语翻译器
type EnglishTranslator struct{}

func (*EnglishTranslator) Translate() string {
	return "这是英语翻译器"
}

//日语翻译器
type JapaneseTranslator struct{}

func (*JapaneseTranslator) Translate() string {
	return "这是日语翻译器"
}
