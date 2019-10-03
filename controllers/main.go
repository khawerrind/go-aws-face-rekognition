package controllers

import (
	"fmt"
	"io/ioutil"

	"github.com/gin-gonic/gin"
	"github.com/khawerrind/go-aws-face-rekognition/services/aws"
	"github.com/khawerrind/go-aws-face-rekognition/services/envvar"
)

type MainController struct{}

func (main *MainController) CompareFaces(c *gin.Context) {
	file, _, err := c.Request.FormFile("file")

	if err != nil {
		c.JSON(406, gin.H{"message": err.Error()})
		c.Abort()
		return
	}

	defer file.Close()

	r, err := ioutil.ReadAll(file)
	if err != nil {
		c.JSON(406, gin.H{"message": err.Error()})
		c.Abort()
		return
	}

	folderPath := c.PostForm("folder_path")

	result, err := aws.CompareFaces(folderPath, r)
	if err != nil {
		c.JSON(406, gin.H{"message": err.Error()})
		c.Abort()
		return
	}

	finalRes := []string{}
	for _, res := range result {
		if !res.Error {
			finalRes = append(finalRes,
				fmt.Sprintf("https://%s.s3.amazonaws.com/%s", envvar.GetEnv(envvar.AWS_S3_BUCKET_KEY), res.Key))
		}
	}

	c.JSON(200, gin.H{"result": finalRes})
}
