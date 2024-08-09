package main

import "github.com/gin-gonic/gin"

func allRoutes(route *gin.Engine) {

	health := route.Group("/health")
	health.GET("/ready", healthhandler) //this is for the server health check

	user := route.Group("/user")
	user.POST("/signUp", signUpUserHandler) // this handler is to create a user

}
