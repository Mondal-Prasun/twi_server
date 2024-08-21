package main

import "github.com/gin-gonic/gin"

func allRoutes(route *gin.Engine, apiCfg *apiCfg) {

	health := route.Group("/health")
	health.GET("/ready", healthhandler) //this is for the server health check

	user := route.Group("/user")
	user.POST("/signUp", apiCfg.signUpUserHandler) // this handler is to create a user
	user.POST("/logIn", apiCfg.logInUserHandler)   // this handler is to login the user

	post := route.Group("/post")
	post.POST("/createPost", convertImage(), apiCfg.createPostHandler) //this handler is to post a feed
	post.DELETE("/deletePost/:postId", apiCfg.deletePostHandler)       // this handler is to delete a post
	post.PATCH("/likePost/:postId", apiCfg.likePostHandler)            //this handler is to like a post
	post.POST("/commentPost", apiCfg.commentPostHandler)               //this handler is to comment to post
	post.GET("/getComments/:postId", apiCfg.getAllComments)            //this handler is to get all comments of that post
}
