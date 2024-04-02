package router

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type QuestionText struct {
	Text string `json:"text"`
}

type Questions struct {
	Category string `json:"category"`
	Id string `json:"id"`
	Tags []string `json:"tags"`
	Difficulty string `json:"difficulty"`
	Regions []string `json:"regions"`
	IsNiche bool `json:"isNiche"`
	Question QuestionText `json:"question"`
	CorrectAnswer string `json:"correctAnswer"`
	IncorrectAnswers []string `json:"incorrectAnswers"`
	Type string `json:"type"` 
} 

func getQuestions(c *gin.Context){
	resp, err := http.Get("https://the-trivia-api.com/v2/questions?limit=50")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch questions"})
		return
	}
	defer resp.Body.Close()
	
	if resp.StatusCode != http.StatusOK {
		c.JSON(resp.StatusCode, gin.H{"error": "Failed to fetch questions"})
		return
	}

	responseData, err := io.ReadAll(resp.Body)
    if err != nil {
        fmt.Print(err)
    }

	var result []Questions 
	 
	err = json.Unmarshal(responseData, &result)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse response"})
		return
	}
	
	c.IndentedJSON(http.StatusOK, result)
}

func InitRouter() *gin.Engine {
	router := gin.Default()

	router.Use(cors.Default())
	router.GET("/questions", getQuestions)
	return router
}