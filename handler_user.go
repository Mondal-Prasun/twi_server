package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

var users = []User{}

type User struct {
	Id          uuid.UUID `json:"id"`
	Username    string    `json:"username"`
	Password    string    `json:"password"`
	Email       string    `json:"email"`
	AccessToken uuid.UUID `json:"accessToken"`
}

func signUpUserHandler(c *gin.Context) {

	cred := User{}

	if err := c.BindJSON(&cred); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid credential",
		})
		return
	}

	for _, i := range users {
		if i.Username == cred.Username {
			c.JSON(http.StatusForbidden, gin.H{
				"error": "User already excits",
			})
			return
		}
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(cred.Password), bcrypt.DefaultCost)

	if err != nil {

		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	}
	cred.Id = uuid.New()
	cred.AccessToken = uuid.New()
	cred.Password = string(hashedPassword)

	users = append(users, cred)

	c.JSON(http.StatusCreated, users)

}
