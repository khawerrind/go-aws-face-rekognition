package main

import (
	"github.com/gin-gonic/gin"
	"github.com/khawerrind/go-aws-face-rekognition/controllers"
	"github.com/khawerrind/go-aws-face-rekognition/services/envvar"
	"net/http"
)

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
		v1.POST("/findFaces", mainController.FindFaces)
	}

	router.NoRoute(func(c *gin.Context) {
		c.String(http.StatusNotFound, "Not Found")
	})

	router.Run()
}
