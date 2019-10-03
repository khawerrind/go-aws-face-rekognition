package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/khawerrind/go-aws-face-rekognition/controllers"
	"github.com/khawerrind/go-aws-face-rekognition/services/envvar"
)

func init() {
	// loads values from .env into the system
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

func main() {
	//make sure env variables are setup correctly
	envvar.MustGetEnv(envvar.AWS_REGION)
	envvar.MustGetEnv(envvar.AWS_ACCESS_KEY_ID)
	envvar.MustGetEnv(envvar.AWS_SECRET_ACCESS_KEY)
	envvar.MustGetEnv(envvar.AWS_S3_BUCKET_KEY)

	router := gin.Default()

	v1 := router.Group("/v1")
	{
		mainController := new(controllers.MainController)
		v1.POST("/compareFaces", mainController.CompareFaces)
	}

	router.NoRoute(func(c *gin.Context) {
		c.String(http.StatusNotFound, "Not Found")
	})

	router.Run()
}
