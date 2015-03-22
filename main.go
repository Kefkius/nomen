package main

import (
	"fmt"
	"github.com/kefkius/nomen/server"
)

func main() {
	fmt.Println("Server starting")
	userName := "rpcuser"
	userPass := "rpcpassword"
	userHost := "host:port"
	server.Init(userName, userPass, userHost)
}
