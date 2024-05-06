package router

import (
	"fmt"
	"funquizbe/models"
	"net/http"

	"github.com/gin-gonic/gin"
)


func postQuestion(c *gin.Context) {
	var newQuestion models.UserQuestion
	if err := c.ShouldBindJSON(&newQuestion); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON data"})
        return
    }

	fmt.Println(newQuestion)
	c.JSON(http.StatusOK, newQuestion)
}