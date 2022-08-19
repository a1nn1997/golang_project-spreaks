package main

import (
	"log"
	"os"

	routes "github.com/a1nn1997/golang-jwt/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil{
		log.Fatal("error in loading env file")
	}
	port :=os.Getenv("PORT")
	if port == "" {
		port="8081"
	}
	router := gin.New()
	router.Use(gin.Logger())
	routes.AuthRoutes(router)
	routes.UserRoutes(router)
	
	router.GET("/api-1", func(c *gin.Context){
		c.JSON(200, gin.H{"success":"Acess granted for api-1"})
	})

	router.GET("/api-2", func(c *gin.Context){
		c.JSON(200, gin.H{"success":"Acess granted for api-2"})
	})
	router.Run(":"+port)
}