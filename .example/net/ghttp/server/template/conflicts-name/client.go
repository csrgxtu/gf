package main

import (
	"github.com/csrgxtu/gf/frame/g"
	"github.com/csrgxtu/gf/net/ghttp"
)

// https://github.com/csrgxtu/gf/issues/437
func main() {
	s := g.Server()
	s.BindHandler("/", func(r *ghttp.Request) {
		r.Response.WriteTpl("client/layout.html")
	})
	s.SetPort(8199)
	s.Run()
}
