package router

import (
	"funquizbe/db"
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

    db, err := db.ConnectDB()
	if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Database connection error"})
        return
    }

    sqlStatement := `
    INSERT INTO questions (question, correctanswer, wronganswer1, wronganswer2, wronganswer3)
    VALUES ($1, $2, $3, $4, $5)
    RETURNING id`

    stmt, err := db.Prepare(sqlStatement)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Database preparation error"})
        return
    }
    defer stmt.Close()

    _, err = stmt.Exec(newQuestion.Question, newQuestion.CorrectAnswer, newQuestion.WrongAnswer1, newQuestion.WrongAnswer2, newQuestion.WrongAnswer3)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Database insertion error"})
        return
    }
	defer db.Close()

    c.JSON(http.StatusCreated, newQuestion)
}