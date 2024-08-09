package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/Mondal-Prasun/custom_backend/internal/database"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func (apiCfg *apiCfg) signUpUserHandler(c *gin.Context) {

	cred := User{}

	if err := c.BindJSON(&cred); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid credential",
		})
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(cred.Password), bcrypt.DefaultCost)

	if err != nil {

		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	cratedUser, err := apiCfg.db.CreateUser(context.Background(), database.CreateUserParams{
		ID:       uuid.New(),
		Username: cred.Username,
		Password: string(hashedPassword),
		Image: sql.NullString{
			Valid: false,
		},
		Email:       cred.Email,
		Createdat:   time.Now().Local(),
		Updatedat:   time.Now().Local(),
		Accesstoken: uuid.New(),
	})

	if err != nil {

		log.Println(err.Error())

		if strings.Contains(err.Error(), "users_email_key") {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "An account already excists with this email accont",
			})
			return
		} else if strings.Contains(err.Error(), "users_username_key") {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Username already excists",
			})
			return
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": fmt.Sprintf("Internal server error: %v", err.Error()),
			})
			return
		}

	}

	c.JSON(http.StatusCreated, gin.H{
		"message":     "user created",
		"id":          cratedUser.ID,
		"email":       cratedUser.Email,
		"image":       cratedUser.Image,
		"accessToken": cratedUser.Accesstoken.String(),
	})

}
