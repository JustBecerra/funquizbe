package main

import (
	router "funquizbe/routes"
)

func main(){
 router := router.InitRouter()
 router.Run(":8080")
}