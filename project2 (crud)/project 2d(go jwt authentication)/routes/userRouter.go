package routes

import(
	controller "github.com/a1nn1997/golang-jwt/controller"
	"github.com/a1nn1997/golang-jwt/middleware"
	"github.com/gin-gonic/gin"

)
func UserRoutes(incomingRoutes *gin.Engine){
	incomingRoutes.Use(middleware.Authenticate())
	incomingRoutes.GET("/users/all",controller.GetUsers())
	incomingRoutes.GET("/users/:user_id",controller.GetUserById())
}