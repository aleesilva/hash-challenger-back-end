package main

import (
	ginServer "github.com/aleesilva/hash-challenger-back-end/config"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load(".env")
	server := ginServer.MakeServer()
	server.Run()
}
