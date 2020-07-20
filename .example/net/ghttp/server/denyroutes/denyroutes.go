package main

import "github.com/csrgxtu/gf/frame/g"

func main() {
	s := g.Server()
	s.SetDenyRoutes([]string{
		"/config*",
	})
	s.SetPort(8299)
	s.Run()
}
