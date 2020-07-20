package main

import (
	"github.com/csrgxtu/gf/frame/g"
	"github.com/csrgxtu/gf/net/ghttp"
)

func main() {
	s := g.Server()
	s.SetNameToUriType(ghttp.URI_TYPE_FULLNAME)
	s.EnableAdmin()
	s.BindHandler("/", func(r *ghttp.Request) {
		r.Response.Write("hello world")
	})
	s.SetPort(8199)
	s.Run()
}
