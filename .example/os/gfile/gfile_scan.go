package main

import (
	"github.com/csrgxtu/gf/os/gfile"
	"github.com/csrgxtu/gf/util/gutil"
)

func main() {
	gutil.Dump(gfile.ScanDir("/Users/john/Documents", "*.*"))
	gutil.Dump(gfile.ScanDir("/home/john/temp/newproject", "*", true))
}
