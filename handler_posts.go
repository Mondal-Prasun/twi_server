package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"time"

	"github.com/Mondal-Prasun/custom_backend/internal/database"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (apiCfg *apiCfg) createPostHandler(c *gin.Context) {

	imageBse64 := c.MustGet("imageBase64").(string)
	// imageName := c.MustGet("imageName")
	// imageSize := c.MustGet("imageSize")

	userData := c.PostForm("data")

	if userData == "" {
		c.JSON(404, gin.H{
			"error": "Userdata is required",
		})
		return
	}

	data := struct {
		Username    string    `json:"username"`
		UserId      uuid.UUID `json:"userId"`
		Contenttext string    `json:"contentText"`
	}{}

	if err := json.Unmarshal([]byte(userData), &data); err != nil {
		c.JSON(401, gin.H{
			"error": err.Error(),
		})
		return
	}

	createdPost, err := apiCfg.db.CreatePost(context.Background(), database.CreatePostParams{
		ID:          uuid.New(),
		Username:    data.Username,
		Userid:      data.UserId,
		Contenttext: data.Contenttext,
		Contextimage: sql.NullString{
			String: string(imageBse64),
			Valid:  true,
		},
		Likes:     0,
		Createdat: time.Now().Local(),
		Updatedat: time.Now().Local(),
	})

	if err != nil {
		c.JSON(501, gin.H{
			"error": fmt.Sprintf("can't create post: %v", err.Error()),
		})
		return
	}

	c.JSON(200, convertPostToSendJson(&createdPost))

}
