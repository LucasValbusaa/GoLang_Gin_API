package main

import (
	"github.com/lucasvalbusa/api-go-gin/database"
	"github.com/lucasvalbusa/api-go-gin/routes"
)

func main() {
	database.ConectaComBancoDeDados()
	routes.HandleRequest()
}
