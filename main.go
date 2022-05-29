package main

import (
	"fmt"
	"rest-api-go/database"
	"rest-api-go/routes"
)

func main() {
	database.ConectaComBancoDeDados()
	fmt.Println("Iniciando o servidor Rest com Go")
	routes.HandleRequest()
}
