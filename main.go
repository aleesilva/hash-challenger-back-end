package main

import server "github.com/aleesilva/hash-challenger-back-end/config"

func main() {
	server := server.MakeServer()
	server.Run()
}
