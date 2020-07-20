package main

import (
	"github.com/csrgxtu/gf/frame/g"
	"github.com/csrgxtu/gf/os/gbuild"
)

func main() {
	g.Dump(gbuild.Info())
	g.Dump(gbuild.Map())
}
