package main

import (
	router "funquizbe/routes"

	"github.com/gin-gonic/gin"
)

// Handler is the exported function Vercel will use as the entry point
func main() {
	gin.SetMode(gin.ReleaseMode)
	r := router.InitRouter()
	r.Run(":8080")
}
