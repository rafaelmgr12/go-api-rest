package routes

import (
	"log"
	"net/http"

	"go-api-rest/controllers"
)

func HandleRequests() {
	http.HandleFunc("/", controllers.Home)
	http.HandleFunc("/api/personalities", controllers.AllPersonalities)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
