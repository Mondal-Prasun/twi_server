package main

import (
	"context"
	"fmt"
	"time"

	"github.com/Mondal-Prasun/custom_backend/internal/database"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// this handler is like a post

func (apiCfg *apiCfg) likePostHandler(c *gin.Context) {
	postId := c.Param("postId")

	postUuid, err := uuid.Parse(postId)

	if err != nil {
		c.JSON(404, gin.H{
			"error": "Invalid post id",
		})
		return
	}

	if err := apiCfg.db.LikePost(context.Background(), database.LikePostParams{
		ID:        postUuid,
		Updatedat: time.Now().Local(),
	}); err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"msg": "post liked",
	})

}

//this is for comment to a post

func (apiCfg *apiCfg) commentPostHandler(c *gin.Context) {
	data := struct {
		PostId  string `json:"postId"`
		UserId  string `json:"userId"`
		Comment string `json:"comment"`
	}{}

	if err := c.BindJSON(&data); err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	postUuid, err := uuid.Parse(data.PostId)

	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	userUuid, err := uuid.Parse(data.UserId)
	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
	}

	if data.Comment == "" {
		c.JSON(403, gin.H{
			"error": "Comment cant be empty",
		})
		return
	}

	postCmt, err := apiCfg.db.GiveComment(context.Background(), database.GiveCommentParams{
		ID:        uuid.New(),
		Comment:   data.Comment,
		Postid:    postUuid,
		Userid:    userUuid,
		Createdat: time.Now().Local(),
		Updatedat: time.Now().Local(),
	})

	if err != nil {
		c.JSON(501, gin.H{
			"error": fmt.Sprintf("Internal server error %v", err.Error()),
		})
		return
	}

	c.JSON(200, gin.H{
		"id":      postCmt.ID,
		"comment": postCmt.Comment,
		"userId":  postCmt.Userid,
	})

}

///this handler is for get all comments

func (apiCfg *apiCfg) getAllComments(c *gin.Context) {
	postId := c.Param("postId")

	postUuid, err := uuid.Parse(postId)

	if err != nil {
		c.JSON(403, gin.H{
			"error": fmt.Sprintf("Invalid post id : %v", err.Error()),
		})
		return
	}

	dBpostComments, err := apiCfg.db.GetAllComment(context.Background(), postUuid)

	if err != nil {
		c.JSON(502, gin.H{
			"error": err.Error(),
		})
		return
	}

	if len(dBpostComments) == 0 {
		c.JSON(203, gin.H{
			"msg": "no comment found",
		})
		return
	}

	postComments := []Comment{}

	for _, cmt := range dBpostComments {
		postComments = append(postComments, convertCommnetToCommentJson(&cmt))
	}

	c.JSON(200, postComments)

}
