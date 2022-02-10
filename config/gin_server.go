package config

import (
	"log"

	"github.com/aleesilva/hash-challenger-back-end/config/routes"
	"github.com/gin-gonic/gin"
)

type GinServer struct {
	port   string
	server *gin.Engine
}

func MakeServer() GinServer {

	return GinServer{
		port:   LoadDotEnvVariable("APP_PORT"),
		server: gin.Default(),
	}
}

func (gs *GinServer) Run() {
	route := routes.Routes(gs.server)

	log.Fatal(route.Run(":" + gs.port))
}
