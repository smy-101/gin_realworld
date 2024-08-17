package main

import (
	"gin_realworld/server"
	_ "gin_realworld/storage"
)

func main() {
	server.RunHttpServer()
}
