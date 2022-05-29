package routes

import (
	"log"
	"net/http"
	"rest-api-go/controllers"

	"github.com/gorilla/mux"
)

func HandleRequest() {
	r := mux.NewRouter()
	r.HandleFunc("/", controllers.Home)
	r.HandleFunc("/Personalidades", controllers.TodasPersonalidades).Methods("Get")
	r.HandleFunc("/Personalidades/{id}", controllers.RetornaUmaPersonalidade).Methods("Get")
	log.Fatal(http.ListenAndServe(":8000", r))
}
