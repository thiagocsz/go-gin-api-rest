package main

import (
	"github.com/thiagocsz/go-gin-api-rest.git/database"
	"github.com/thiagocsz/go-gin-api-rest.git/routes"
)

func main() {
	database.ConectaDB()
	routes.HandleRequests()
}
