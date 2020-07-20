package main

import (
	"fmt"
	"github.com/csrgxtu/gf/frame/g"
	"github.com/csrgxtu/gf/util/gvalid"
)

func main() {
	g.I18n().SetLanguage("cn")
	err := gvalid.Check("", "required", nil)
	fmt.Println(err.String())
}
