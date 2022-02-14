package main

import ginServer "github.com/aleesilva/hash-challenger-back-end/config"

func main() {
	server := ginServer.MakeServer()
	server.Run()
}
