package main

import (
	"fmt"
	"rest-api-go/models"
	"rest-api-go/routes"
)

func main() {
	models.Personalidades = []models.Personalidade{
		{Id: 1, Nome: "Jader", Historia: "História de Badaroska"},
		{Id: 2, Nome: "Juscelino", Historia: "Conto de Campos"},
	}

	fmt.Println("Iniciando o servidor Rest com Go")
	routes.HandleRequest()
}
