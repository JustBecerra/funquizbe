package router

import (
	"fmt"
	"funquizbe/db"
	"funquizbe/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func getUserQuestions(c *gin.Context) {
    var userQuestions []models.UserQuestion

    db, err := db.ConnectDB()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Database connection error"})
        return
    }
    defer db.Close()

    // Prepare the SQL statement
    sqlStatement := `
        SELECT * FROM questions`

    // Execute the SQL statement
    rows, err := db.Query(sqlStatement)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Database query error"})
        return
    }
    defer rows.Close()

    // Iterate through the result set
    for rows.Next() {
        var userQuestion models.UserQuestion
        // Scan each row into a UserQuestion struct
        err := rows.Scan(&userQuestion.ID, &userQuestion.Question, &userQuestion.CorrectAnswer, &userQuestion.IncorrectAnswers)
        if err != nil {
			fmt.Println("Error scanning row:", err)
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Error scanning rows"})
            return
        }
        // Append the UserQuestion to the slice
        userQuestions = append(userQuestions, userQuestion)
    }
    // Check for errors during row iteration
    if err := rows.Err(); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Error iterating over rows"})
        return
    }

    c.JSON(http.StatusOK, userQuestions)
}