package main

import (
	"encoding/base64"
	"io"

	"github.com/gin-gonic/gin"
)

func convertImage() gin.HandlerFunc {
	return func(c *gin.Context) {
		imageFile, err := c.FormFile("image")

		if err != nil {
			c.JSON(404, "File not found")
			return
		}

		openedImage, err := imageFile.Open()
		if err != nil {
			c.JSON(501, gin.H{
				"error": err.Error(),
			})
			return
		}

		defer openedImage.Close()

		imageBytes, err := io.ReadAll(openedImage)

		if err != nil {
			c.JSON(503, gin.H{
				"error": "cant open file bytes",
			})
			return
		}

		imageBase64 := base64.StdEncoding.EncodeToString(imageBytes)

		c.Set("imageBase64", imageBase64)
		c.Set("imageName", imageFile.Filename)
		c.Set("imageSize", imageFile.Size)
		c.Next()

	}
}
