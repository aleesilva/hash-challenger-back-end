package routes

import (
	"github.com/aleesilva/hash-challenger-back-end/controllers"
	"github.com/gin-gonic/gin"
)

func Routes(router *gin.Engine) *gin.Engine {
	api := router.Group("api")
	{
		cart := api.Group("cart")
		{
			cart.GET("checkout", controllers.Checkout)
		}
	}

	return router
}
