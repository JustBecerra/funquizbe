package main

import (
	router "funquizbe/routes"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Handler is the exported function Vercel will use as the entry point
func Handler(w http.ResponseWriter, r *http.Request) {
	// Initialize the Gin router
	router := router.InitRouter()

	// Set Gin to release mode
	gin.SetMode(gin.ReleaseMode)

	// Serve the request using Gin
	router.ServeHTTP(w, r)
}
