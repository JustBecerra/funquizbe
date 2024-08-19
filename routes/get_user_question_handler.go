package router

import (
	"database/sql"
	"funquizbe/db"
	"funquizbe/models"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// Helper function to convert PostgreSQL array format to a slice of strings
func parsePostgresArray(raw sql.RawBytes) ([]string, error) {
	// Convert raw bytes to string
	str := string(raw)
	// Remove curly braces and whitespace
	str = strings.Trim(str, "{}")
	if str == "" {
		return []string{}, nil
	}
	// Split by commas
	elements := strings.Split(str, ",")
	// Trim quotes from elements if necessary
	for i := range elements {
		elements[i] = strings.Trim(elements[i], `"`)
	}
	return elements, nil
}

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
		SELECT id, question, correctanswer, incorrectanswers
		FROM questions`

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
		var incorrectAnswersRaw sql.RawBytes

		// Scan each row into a UserQuestion struct
		err := rows.Scan(&userQuestion.ID, &userQuestion.Question, &userQuestion.CorrectAnswer, &incorrectAnswersRaw)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error scanning rows", "details": err.Error()})
			return
		}

		// Convert the raw byte slice into a slice of strings
		if incorrectAnswersRaw != nil {
			incorrectAnswers, err := parsePostgresArray(incorrectAnswersRaw)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Error parsing incorrect answers", "details": err.Error()})
				return
			}
			userQuestion.IncorrectAnswers = incorrectAnswers
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
