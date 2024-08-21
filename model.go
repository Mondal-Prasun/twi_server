package main

import (
	"time"

	"github.com/Mondal-Prasun/custom_backend/internal/database"
	"github.com/google/uuid"
)

type User struct {
	Id        uuid.UUID `json:"id"`
	Username  string    `json:"username"`
	Password  string    `json:"password"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type Post struct {
	Id           uuid.UUID `json:"id"`
	Username     string    `json:"username"`
	UserId       uuid.UUID `json:"userId"`
	ContentText  string    `json:"contentText"`
	ContentImage string    `json:"contentImage"`
	Likes        int       `json:"likes"`
	CreatedAt    time.Time `json:"createdAt"`
	UpdatedAt    time.Time `json:"updatedAt"`
}

func convertPostToSendJson(p *database.Post) Post {
	return Post{
		Id:           p.ID,
		Username:     p.Username,
		UserId:       p.Userid,
		ContentText:  p.Contenttext,
		ContentImage: p.Contextimage.String,
		Likes:        int(p.Likes),
		CreatedAt:    p.Createdat,
		UpdatedAt:    p.Updatedat,
	}
}

type Comment struct {
	Id        uuid.UUID `json:"id"`
	PostId    uuid.UUID `json:"postId"`
	UserId    uuid.UUID `json:"userId"`
	Comment   string    `json:"comment"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func convertCommnetToCommentJson(c *database.Comment) Comment {
	return Comment{
		Id:        c.ID,
		PostId:    c.Postid,
		UserId:    c.Userid,
		Comment:   c.Comment,
		UpdatedAt: c.Updatedat,
		CreatedAt: c.Createdat,
	}
}
