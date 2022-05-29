package main

import (
	"fmt"
	"rest-api-go/models"
	"rest-api-go/routes"
)

func main() {
	models.Personalidades = []models.Personalidade{
		{Nome: "Jader", Historia: "Badaroska"},
		{Nome: "Juscelino", Historia: "Campos"},
	}

	fmt.Println("Iniciando o servidor Rest com Go")
	routes.HandleRequest()
}
