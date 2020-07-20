package main

import (
	"github.com/csrgxtu/gf/frame/g"
	"github.com/csrgxtu/gf/net/ghttp"
	"github.com/csrgxtu/gf/os/glog"
)

func main() {
	s := g.Server()
	s.SetIndexFolder(true)
	s.BindHandler("/", func(r *ghttp.Request) {
		glog.Println(r.Header)
		r.Response.Write("hello world")
	})
	s.SetPort(8999)
	s.Run()
}
