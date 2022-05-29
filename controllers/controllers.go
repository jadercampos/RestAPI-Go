package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"rest-api-go/models"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<body><h1>Hello World</h1></body>")
}

func TodasPersonalidades(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(models.Personalidades)
}
