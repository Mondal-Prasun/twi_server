package main

import "github.com/gin-gonic/gin"

func allRoutes(route *gin.Engine, apiCfg *apiCfg) {

	health := route.Group("/health")
	health.GET("/ready", healthhandler) //this is for the server health check

	user := route.Group("/user")
	user.POST("/signUp", apiCfg.signUpUserHandler) // this handler is to create a user
	user.POST("/logIn", apiCfg.logInUserHandler)   // this handler is to login the user

	post := route.Group("/post")
	post.POST("/createPost", convertImage(), apiCfg.createPostHandler)

}
