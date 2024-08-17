package router

import (
	"funquizbe/db"
	"net/http"

	"github.com/gin-gonic/gin"
)

func DeleteUserQuestions(c *gin.Context) {
    id := c.Query("id")

    db, err := db.ConnectDB()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Database connection error"})
        return
    }
    defer db.Close()

    sqlStatement := `
    DELETE FROM questions
    WHERE id = $1`

    stmt, err := db.Prepare(sqlStatement)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Database preparation error"})
        return
    }
    defer stmt.Close()

    res, err := stmt.Exec(id)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Database execution error"})
        return
    }

    rowsAffected, err := res.RowsAffected()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching result"})
        return
    }

    if rowsAffected == 0 {
        c.JSON(http.StatusNotFound, gin.H{"error": "Question not found"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Question deleted"})
}

