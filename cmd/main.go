package main

import (
	"github.com/x14n/go-chat-x14n/initliza"
	"github.com/x14n/go-chat-x14n/server"
)

func main() {
	initliza.InitServer()
	server.SCore()

}
