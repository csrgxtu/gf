package main

import (
	"time"

	"github.com/csrgxtu/gf/os/glog"
	"github.com/csrgxtu/gf/os/gtime"
	"github.com/csrgxtu/gf/os/gtimer"
)

func main() {
	gtimer.SetTimeout(3*time.Second, func() {
		glog.SetDebug(false)
	})
	for {
		glog.Debug(gtime.Datetime())
		time.Sleep(time.Second)
	}
}
