package router

import (
	"encoding/json"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func getQuestions(c *gin.Context){
	resp, err := http.Get("https://the-trivia-api.com/v2/questions")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch questions"})
		return
	}
	defer resp.Body.Close()
	
	if resp.StatusCode != http.StatusOK {
		c.JSON(resp.StatusCode, gin.H{"error": "Failed to fetch questions"})
		return
	}
	
	var result interface{} // Change the type to match the expected response
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse response"})
		return
	}
	
	// Return the result to the client
	c.JSON(http.StatusOK, result)
}

func InitRouter() *gin.Engine {
	router := gin.Default()

	router.Use(cors.Default())
	router.GET("/questions", getQuestions)
	return router
}