package main

import (
	"github.com/csrgxtu/gf/frame/g"
)

func main() {
	db := g.DB()
	db.SetDebug(true)

	tables, err := db.Tables()
	if err != nil {
		panic(err)
	}
	if tables != nil {
		g.Dump(tables)
	}
}
