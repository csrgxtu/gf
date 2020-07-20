package main

import (
	"github.com/csrgxtu/gf/frame/g"
	"github.com/csrgxtu/gf/util/gvalid"
)

// string默认值校验
func main() {
	type User struct {
		Uid string `gvalid:"uid@integer"`
	}

	user := &User{}

	g.Dump(gvalid.CheckStruct(user, nil))
}
