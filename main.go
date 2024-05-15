package main

import (
	"API/controllers"
	"API/initializers"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func init() {
	initializers.ConnectToDB()
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Starting server on %s:%s...\n", os.Getenv("HOST"), port)
	err = http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
	
	r := gin.Default()
	r.GET("/ping", gin.HandlerFunc(func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "ney",
		})
	}))

	r.PUT("/posts/:id", controllers.PostsUpdate)
	r.GET("/posts", controllers.PostsIndex)
	r.GET("/posts/:id", controllers.PostsShow)
	r.DELETE("/posts/:id", controllers.PostsDelete)
	r.Run() // listen and serve on 0.0.0.0:8080

}
