package main

import (
	"github.com/csrgxtu/gf/frame/g"
)

func main() {
	s := g.Server()
	s.EnableAdmin()
	s.SetPort(8199)
	s.Run()
}
