package router

import (
	"funquizbe/db"
	"funquizbe/models"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// Helper function to convert a slice of strings to a PostgreSQL array string
func toPostgresArray(arr []string) string {
	if len(arr) == 0 {
		return "{}"
	}
	return "{" + strings.Join(arr, ",") + "}"
}

func postQuestion(c *gin.Context) {
	var newQuestion models.UserQuestion
    if err := c.ShouldBindJSON(&newQuestion); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON data", "details": err.Error()})
        return
    }

    db, err := db.ConnectDB()
	if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Database connection error", "details": err.Error()})
        return
    }
    defer db.Close()

    sqlStatement := `
    INSERT INTO questions (question, correctanswer, incorrectanswers)
    VALUES ($1, $2, $3)
    RETURNING id`

    incorrectAnswersArray := toPostgresArray(newQuestion.IncorrectAnswers)

    var id int
    err = db.QueryRow(sqlStatement, newQuestion.Question, newQuestion.CorrectAnswer, incorrectAnswersArray).Scan(&id)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Database insertion error", "details": err.Error()})
        return
    }

    newQuestion.ID = id
    c.JSON(http.StatusCreated, newQuestion)
}
