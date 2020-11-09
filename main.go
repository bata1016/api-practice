package main

import (
	"github.com/bata1016/api-practice/db"
	"github.com/bata1016/api-practice/server"
)

func main() {

	db.Init()
	server.Init()
	db.Close()
}
