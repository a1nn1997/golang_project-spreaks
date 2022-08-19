package routes

import(
	controller "github.com/a1nn1997/golang-jwt/controller"
	"github.com/gin-gonic/gin"

)
func AuthRoutes(incomingRoutes *gin.Engine){
	incomingRoutes.POST("/users/signup",controller.SignUp())
	incomingRoutes.POST("/users/login",controller.Login())
}