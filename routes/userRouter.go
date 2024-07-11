package routes

import (
	"github.com/gin-gonic/gin"
	controllers "github.com/jopaleti/golang-authentication-jwt/controllers"
	"github.com/jopaleti/golang-authentication-jwt/middleware"
)

func UserRoutes(incomingRoutes *gin.Engine) {
	// Middleware to ensure user has been authenticated before accessing this content
	// Also known as AUTHORISATION
	incomingRoutes.Use(middleware.Authenticate())
	incomingRoutes.GET("/users", controllers.GetUsers())
	incomingRoutes.GET("/users/:user_id", controllers.GetUser())
}