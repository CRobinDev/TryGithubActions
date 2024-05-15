package main

import (
	"API/controllers"
	"API/initializers"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	initializers.ConnectToDB()

	r := gin.Default()
	r.GET("/ping", gin.HandlerFunc(func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "hai",
		})
	}))

	r.PUT("/posts/:id", controllers.PostsUpdate)
	r.GET("/posts", controllers.PostsIndex)
	r.GET("/posts/:id", controllers.PostsShow)
	r.DELETE("/posts/:id", controllers.PostsDelete)
	r.Run()

}
