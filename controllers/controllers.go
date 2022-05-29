package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"rest-api-go/models"
	"strconv"

	"github.com/gorilla/mux"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<body><h1>Hello World</h1></body>")
}

func TodasPersonalidades(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(models.Personalidades)
}
func RetornaUmaPersonalidade(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	for _, personalidade := range models.Personalidades {
		if strconv.Itoa(personalidade.Id) == id {
			json.NewEncoder(w).Encode(personalidade)
		}
	}

}
