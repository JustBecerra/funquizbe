package router

import (
	"encoding/json"
	"fmt"
	"funquizbe/models"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

func getQuestions(c *gin.Context) {
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

	var result []models.Question

	err = json.Unmarshal(responseData, &result)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse response"})
		return
	}

	c.IndentedJSON(http.StatusOK, result)
}