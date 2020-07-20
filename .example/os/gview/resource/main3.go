package main

import (
	"fmt"

	"github.com/csrgxtu/gf/frame/g"
	"github.com/csrgxtu/gf/os/gres"
	_ "github.com/csrgxtu/gf/os/gres/testdata"
)

func main() {
	gres.Dump()

	v := g.View()
	s, err := v.Parse("index.html")
	fmt.Println(err)
	fmt.Println(s)
}
