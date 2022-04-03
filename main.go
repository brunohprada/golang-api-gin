package main

import (
	"github.com/brunohprada/go-api-gin/database"
	"github.com/brunohprada/go-api-gin/routes"
)

func main() {
	database.ConectaComBancoDeDados()
	routes.HandleRequests()
}
