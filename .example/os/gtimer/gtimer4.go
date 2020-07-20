package main

import (
	"time"

	"github.com/csrgxtu/gf/os/glog"
	"github.com/csrgxtu/gf/os/gtimer"
)

func main() {
	interval := time.Second
	gtimer.AddTimes(interval, 2, func() {
		glog.Println("doing1")
	})
	gtimer.AddTimes(interval, 2, func() {
		glog.Println("doing2")
	})

	select {}
}
