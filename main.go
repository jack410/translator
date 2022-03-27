package main

import (
	"fmt"
	"translator_factory/factory"
)

func main() {
	var lan int
	fmt.Printf("%s\r\n%s\r\n", "以下是可翻译的语言种类，请输入代表数字", "1：德语、2：英语、3：日语")
	fmt.Scanln(&lan)
	fmt.Println(factory.CreateTranslator(lan).Translate())
}
