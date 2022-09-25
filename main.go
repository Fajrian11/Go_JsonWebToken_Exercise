package main

import (
	"Go_JsonWebToken_Exercise/database"
	"Go_JsonWebToken_Exercise/router"
)

func main() {
	database.StartDB()
	r := router.StartApp()
	r.Run(":3000")
}
