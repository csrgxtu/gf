package main

import (
	"fmt"

	"github.com/csrgxtu/gf/frame/g"
	"github.com/csrgxtu/gf/util/gconv"
)

func main() {
	conn := g.Redis().Conn()
	defer conn.Close()
	conn.Do("SET", "k", "v")
	v, _ := conn.Do("GET", "k")
	fmt.Println(gconv.String(v))
}
