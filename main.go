package main

import (
	"gin-rest-api/database"
	"gin-rest-api/routes"
)

func main() {
	database.Connect()
	routes.HandleRequests()
}
