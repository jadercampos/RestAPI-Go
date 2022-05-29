package routes

import (
	"log"
	"net/http"
	"rest-api-go/controllers"
	"rest-api-go/middleware"

	"github.com/gorilla/mux"
)

func HandleRequest() {
	r := mux.NewRouter()
	r.Use(middleware.ContentTypeMiddleware)
	r.HandleFunc("/", controllers.Home)
	r.HandleFunc("/Personalidades", controllers.TodasPersonalidades).Methods("Get")
	r.HandleFunc("/Personalidades/{id}", controllers.RetornaUmaPersonalidade).Methods("Get")
	r.HandleFunc("/Personalidades", controllers.CriaUmaNovaPersonalidade).Methods("Post")
	r.HandleFunc("/Personalidades/{id}", controllers.DeletaUmaPersonalidade).Methods("Delete")
	r.HandleFunc("/Personalidades/{id}", controllers.EditaPersonalidade).Methods("Put")
	log.Fatal(http.ListenAndServe(":8000", r))
}
