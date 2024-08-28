package main

import (
	router "funquizbe/routes"

	"github.com/gin-gonic/gin"
)

func main(){
 router := router.InitRouter()
 gin.SetMode(gin.ReleaseMode)
 router.Run(":8080")
}