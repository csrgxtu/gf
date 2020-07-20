package main

import (
	"fmt"
	"github.com/csrgxtu/gf/debug/gdebug"
)

func main() {
	gdebug.PrintStack()
	fmt.Println(gdebug.CallerPackage())
	fmt.Println(gdebug.CallerFunction())
}
