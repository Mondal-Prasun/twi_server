package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func healthhandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "server is ready",
	})
}
