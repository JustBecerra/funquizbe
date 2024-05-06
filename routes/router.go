package router

import (
	"fmt"

	"funquizbe/db"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)



func InitRouter() *gin.Engine {
	err := godotenv.Load()
    if err != nil {
        fmt.Println("Error loading .env file")
    }

	db, err := db.ConnectDB()
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()

    // Ping the database to test the connection
    err = db.Ping()
    if err != nil {
        panic(err)
    }

    fmt.Println("Successfully connected to PostgreSQL database!")
	router := gin.Default()

	router.Use(cors.Default())
	router.GET("/questions", getQuestions)
	router.POST("/userquestion", postQuestion)
	return router
}