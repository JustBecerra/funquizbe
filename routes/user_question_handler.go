package router

import (
	"funquizbe/db"
	"funquizbe/models"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
)


func postQuestion(c *gin.Context) {
	var newQuestion models.UserQuestion
	if err := c.BindJSON(&newQuestion); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON data"})
        return
    }

	db := db.GetDB()

	sqlStatement := `
    INSERT INTO questions (question, correctanswer, incorrectanswers)
    VALUES ($1, $2, $3)
    RETURNING id`
	// var questionID int

	// err := db.Prepare(sqlStatement, newQuestion.Question, newQuestion.CorrectAnswer, pq.Array(newQuestion.IncorrectAnswers)).Scan(&questionID)
	// if err != nil {
	// 	c.JSON(http.StatusInternalServerError, gin.H{"error": "Database insertion error"})
	// 	return
	// }
	stmt, err := db.Prepare(sqlStatement)
	if err != nil {
		log.Fatal(err)
	}
	
	if _, err := stmt.Exec(newQuestion.Question, newQuestion.CorrectAnswer, pq.Array(newQuestion.IncorrectAnswers)); err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	c.JSON(http.StatusCreated, newQuestion)

	// c.JSON(http.StatusOK, gin.H{"message": "Question posted successfully"})
}