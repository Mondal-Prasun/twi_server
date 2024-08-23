package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/Mondal-Prasun/custom_backend/internal/database"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type apiCfg struct {
	db *database.Queries
}

func main() {
	godotenv.Load(".env")

	route := gin.Default()

	portNum := os.Getenv("PORT")

	dbportNum := os.Getenv("DB_URL")

	dbConnection, err := sql.Open("postgres", dbportNum)

	if err != nil {
		log.Fatal("Database connection error: ", err)
	}

	db := database.New(dbConnection)

	apiCfg := apiCfg{
		db: db,
	}

	// all the routes should be here
	allRoutes(route, &apiCfg)

	err = route.Run(portNum)

	if err != nil {
		log.Fatal("Problem occoured in server: ", err)
	}

}
