package routes

import (
	"log"
	"net/http"

	"go-api-rest/controllers"

	"github.com/gorilla/mux"
)

func HandleRequests() {
	r := mux.NewRouter()
	r.HandleFunc("/", controllers.Home)
	r.HandleFunc("/personalities", controllers.AllPersonalities).Methods("GET")
	r.HandleFunc("/personalities/{id}", controllers.ReturnOnePersonality).Methods("GET")
	log.Fatal(http.ListenAndServe(":8080", r))
}
