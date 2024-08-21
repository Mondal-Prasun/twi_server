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

//this handler for creating user

func (apiCfg *apiCfg) signUpUserHandler(c *gin.Context) {

	cred := User{}

	if err := c.BindJSON(&cred); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid credential",
		})
		return
	}
	//password hashed here by bcrypt library
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(cred.Password), bcrypt.DefaultCost)

	if err != nil {

		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	// new user created here
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

	//all database err handled here
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
		"username":    cratedUser.Username,
		"email":       cratedUser.Email,
		"image":       cratedUser.Image,
		"accessToken": cratedUser.Accesstoken.String(),
	})

}

// this handler is for login user
func (apiCfg *apiCfg) logInUserHandler(c *gin.Context) {
	cred := struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}{}

	if err := c.BindJSON(&cred); err != nil {
		c.JSON(http.StatusNotAcceptable, gin.H{
			"error": err.Error(),
		})
		return
	}

	if cred.Email == "" || cred.Password == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Email and password is required",
		})
		return
	}

	// this func matches email and return password and id
	dbUser, err := apiCfg.db.GetUserPasswordByEmail(context.Background(), cred.Email)

	if err != nil {
		c.JSON(404, gin.H{
			"error": "No user found",
		})
		return
	}

	passErr := bcrypt.CompareHashAndPassword([]byte(dbUser.Password), []byte(cred.Password))

	if passErr != nil {
		c.JSON(404, gin.H{
			"error": "Invalid password",
		})
		return
	}

	dbUserDetails, err := apiCfg.db.RefreshUserAccessToken(context.Background(), database.RefreshUserAccessTokenParams{
		ID:          dbUser.ID,
		Accesstoken: uuid.New(),
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{
		"id":          dbUserDetails.ID,
		"username":    dbUserDetails.Username,
		"email":       dbUserDetails.Email,
		"image":       dbUserDetails.Image,
		"accessToken": dbUserDetails.Accesstoken,
	})

}

// this handler is for user details

func (apiCfg *apiCfg) getUserDetails(c *gin.Context) {
	userId := c.Param("userId")

	userUuid, err := uuid.Parse(userId)

	if err != nil {
		c.JSON(400, gin.H{
			"error": fmt.Sprintf("Invalid user id %v", err.Error()),
		})
		return
	}

	dbUserDetails, err := apiCfg.db.GetUserDetails(context.Background(), userUuid)

	if err != nil {
		c.JSON(500, gin.H{
			"error": fmt.Sprintf("Invalid user id %v", err.Error()),
		})
		return
	}

	c.JSON(202, gin.H{
		"id":       dbUserDetails.ID,
		"username": dbUserDetails.Username,
		"image":    dbUserDetails.Image,
	})

}
