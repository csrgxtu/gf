package main

import (
	"github.com/csrgxtu/gf/frame/g"
	"github.com/csrgxtu/gf/os/glog"
)

func main() {
	err := glog.SetConfigWithMap(g.Map{
		"prefix": "[TEST]",
	})
	if err != nil {
		panic(err)
	}
	glog.Info(1)
}
