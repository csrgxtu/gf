package main

import (
	"errors"

	"github.com/csrgxtu/gf/os/glog"

	"github.com/csrgxtu/gf/errors/gerror"
)

func Error1() error {
	return errors.New("test1")
}

func Error2() error {
	return gerror.New("test2")
}

func main() {
	glog.Println(Error1())
	glog.Println(Error2())
}
