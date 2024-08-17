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

    err = db.Ping()
    if err != nil {
        panic(err)
    }
	defer db.Close()

	router := gin.Default()

	router.Use(cors.Default())
	router.GET("/questions", getQuestions)
	router.GET("/userquestions", getUserQuestions)
	router.POST("/userquestion", postQuestion)
	router.DELETE("/deleteuserquestion", DeleteUserQuestions)
	return router
}