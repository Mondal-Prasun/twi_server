package main

import (
	"context"
	"fmt"

	"github.com/Mondal-Prasun/custom_backend/internal/database"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// this handler is to follow a user
func (apiCfg *apiCfg) followUserHandler(c *gin.Context) {

	dataIds := struct {
		FollowerId string `json:"followerId"`
		FollowedId string `json:"followedId"`
	}{}

	if err := c.BindJSON(&dataIds); err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	followerUuid, err := uuid.Parse(dataIds.FollowerId)

	if err != nil {
		c.JSON(404, gin.H{
			"error": fmt.Sprintf("FollowerUuid: %v", err.Error()),
		})
		return
	}

	followedUuid, err := uuid.Parse(dataIds.FollowedId)

	if err != nil {
		c.JSON(404, gin.H{
			"error": fmt.Sprintf("FollowedId: %v", err.Error()),
		})
		return
	}

	err = apiCfg.db.FollowUser(context.Background(), database.FollowUserParams{
		Followerid: followerUuid,
		Followedid: followedUuid,
	})

	if err != nil {
		c.JSON(503, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"msg": "user followed",
	})

}

//this handler is to unfollow a user

func (apiCfg *apiCfg) unFollowUserHandler(c *gin.Context) {
	dataIds := struct {
		FollowerId string `json:"followerId"`
		FollowedId string `json:"followedId"`
	}{}

	if err := c.BindJSON(&dataIds); err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	followerUuid, err := uuid.Parse(dataIds.FollowerId)

	if err != nil {
		c.JSON(404, gin.H{
			"error": fmt.Sprintf("FollowerUuid: %v", err.Error()),
		})
		return
	}

	followedUuid, err := uuid.Parse(dataIds.FollowedId)

	if err != nil {
		c.JSON(404, gin.H{
			"error": fmt.Sprintf("FollowedId: %v", err.Error()),
		})
		return
	}

	err = apiCfg.db.UnFollowUser(context.Background(), database.UnFollowUserParams{
		Followerid: followedUuid,
		Followedid: followerUuid,
	})

	if err != nil {
		c.JSON(503, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"msg": "user unfollowed",
	})

}

//this handler is for getting all the followed Ids

func (apiCfg *apiCfg) getAllFollowedIds(c *gin.Context) {

	followerUuid, err := uuid.Parse(c.Param("followerId"))

	if err != nil {
		c.JSON(404, gin.H{
			"error": fmt.Sprintf("FollowerUuid: %v", err.Error()),
		})
		return
	}

	followedUuids, err := apiCfg.db.FollowedIds(context.Background(), followerUuid)

	if err != nil {
		c.JSON(503, gin.H{
			"error": fmt.Sprintf("Internal server error: %v", err.Error()),
		})
		return
	}

	c.JSON(201, followedUuids)

}

//this halder is for getting all the followers

func (apiCfg *apiCfg) getAllFollowerIds(c *gin.Context) {

	followedUuid, err := uuid.Parse(c.Param("followedId"))

	if err != nil {
		c.JSON(404, gin.H{
			"error": fmt.Sprintf("FollowedUuid: %v", err.Error()),
		})
		return
	}

	followerUuids, err := apiCfg.db.FollowerIds(context.Background(), followedUuid)

	if err != nil {
		c.JSON(503, gin.H{
			"error": fmt.Sprintf("Internal server error: %v", err.Error()),
		})
		return
	}

	c.JSON(201, followerUuids)

}

//this hanlder is to get all count followers

func (apiCfg *apiCfg) getAllFollowerCount(c *gin.Context) {

	followedUuid, err := uuid.Parse(c.Param("followedId"))

	if err != nil {
		c.JSON(404, gin.H{
			"error": fmt.Sprintf("FollowedUuid: %v", err.Error()),
		})
		return
	}

	followerCount, err := apiCfg.db.FollowerCount(context.Background(), followedUuid)

	if err != nil {
		c.JSON(503, gin.H{
			"error": fmt.Sprintf("Internal server error: %v", err.Error()),
		})
		return
	}

	c.JSON(201, gin.H{
		"count": followerCount,
	})

}
