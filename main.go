package main

import (
	"funquizbe/router"
)

func main(){
 router := router.InitRouter()
 router.Run(":8080")
}