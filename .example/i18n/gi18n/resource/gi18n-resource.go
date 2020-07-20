package main

import (
	"fmt"

	"github.com/csrgxtu/gf/frame/g"

	_ "github.com/csrgxtu/gf/os/gres/testdata"
)

func main() {
	m := g.I18n()
	m.SetLanguage("ja")
	err := m.SetPath("/i18n-dir")
	if err != nil {
		panic(err)
	}
	fmt.Println(m.Translate(`hello`))
	fmt.Println(m.Translate(`{#hello}{#world}!`))
}
