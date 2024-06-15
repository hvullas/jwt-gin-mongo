package routes

import (
	controller "github.com/hvullas/jwt-gin-mongo/controllers"
	"github.com/hvullas/jwt-gin-mongo/middleware"

	"github.com/gin-gonic/gin"
)

func UserRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.Use(middleware.Authenticate())
	incomingRoutes.GET("/users", controller.GetUsers())
	incomingRoutes.GET("/users/:user_id", controller.GetUser())
}
