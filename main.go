package main

import (
	"log"

	"github.com/gin-gonic/gin"
)

func main() {

	route := gin.Default()

	// all the routes should be here
	allRoutes(route)

	err := route.Run(":8080")

	if err != nil {
		log.Fatal("Problem occoured in server: ", err)
	}

}
