package main

import (
	"fmt"

	"github.com/csrgxtu/gf/i18n/gi18n"
)

func main() {
	t := gi18n.New()
	t.SetPath("/Users/john/Workspace/Go/GOPATH/src/github.com/csrgxtu/gf/.example/i18n/gi18n/i18n")
	t.SetLanguage("en")
	fmt.Println(t.Translate(`hello`))
	fmt.Println(t.Translate(`{#hello}{#world}!`))

	t.SetLanguage("ja")
	fmt.Println(t.Translate(`hello`))
	fmt.Println(t.Translate(`{#hello}{#world}!`))

	t.SetLanguage("ru")
	fmt.Println(t.Translate(`hello`))
	fmt.Println(t.Translate(`{#hello}{#world}!`))

	fmt.Println(t.Translate(`hello`, "zh-CN"))
	fmt.Println(t.Translate(`{#hello}{#world}!`, "zh-CN"))
}
