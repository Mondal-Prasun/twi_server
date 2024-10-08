// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0

package database

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
)

type Comment struct {
	ID        uuid.UUID
	Comment   string
	Postid    uuid.UUID
	Userid    uuid.UUID
	Createdat time.Time
	Updatedat time.Time
}

type Post struct {
	ID           uuid.UUID
	Username     string
	Userid       uuid.UUID
	Contenttext  string
	Contextimage sql.NullString
	Likes        int32
	Createdat    time.Time
	Updatedat    time.Time
}

type User struct {
	ID          uuid.UUID
	Username    string
	Password    string
	Image       sql.NullString
	Email       string
	Createdat   time.Time
	Updatedat   time.Time
	Accesstoken uuid.UUID
}

type UserFollow struct {
	Followerid uuid.UUID
	Followedid uuid.UUID
	Followdate sql.NullTime
}
