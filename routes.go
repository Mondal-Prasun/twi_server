package main

import "github.com/gin-gonic/gin"

func allRoutes(route *gin.Engine, apiCfg *apiCfg) {

	health := route.Group("/health")
	health.GET("/ready", healthhandler) //this is for the server health check

	user := route.Group("/user")
	user.POST("/signUp", apiCfg.signUpUserHandler)                                        // this handler is to create a user
	user.POST("/logIn", apiCfg.logInUserHandler)                                          // this handler is to login the user
	user.GET("/getUserDetails/:userId", apiCfg.getUserDetails)                            // this handler is to get user details
	user.PATCH("/uploadUserImage/:userId", convertImage(), apiCfg.uploadUserProfileImage) // this handeler is for uploading user image

	post := route.Group("/post")
	post.POST("/createPost", convertImage(), apiCfg.createPostHandler) //this handler is to post a feed
	post.DELETE("/deletePost/:postId", apiCfg.deletePostHandler)       //this handler is to delete a post
	post.PATCH("/likePost/:postId", apiCfg.likePostHandler)            //this handler is to like a post
	post.POST("/commentPost", apiCfg.commentPostHandler)               //this handler is to comment to post
	post.GET("/getComments/:postId", apiCfg.getAllComments)            //this handler is to get all comments of that post

	follow := route.Group("/follow")
	follow.POST("/followUser", apiCfg.followUserHandler)                     // this handler is for follow a user
	follow.DELETE("/unFollowUser", apiCfg.unFollowUserHandler)               // this hanlder is for to unfollow a user
	follow.GET("/followedIds/:followerId", apiCfg.getAllFollowedIds)         //this is for getting all followed ids
	follow.GET("/followerIds/:followedId", apiCfg.getAllFollowerIds)         //this is for getting all follower ids
	follow.GET("/followerIds/count/:followedId", apiCfg.getAllFollowerCount) // this is for getting the count of followers id

}
