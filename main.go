package main

import (
	"github.com/igorestevanjasinski/gin-api-rest/database"
	"github.com/igorestevanjasinski/gin-api-rest/routes"
)

func main() {
	database.ConectaComBancoDeDados()
	routes.HandleResquests()
}
