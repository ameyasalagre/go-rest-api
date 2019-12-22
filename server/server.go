package server

import (
	"log"
	"net/http"
	"github.com/go-rest-api/routes"
	"github.com/gorilla/mux"
)

/************************.
*	Start the Server
*	Port 8080
************************/
func Start() {
	router := mux.NewRouter().StrictSlash(true)
	routes.IntiateRoutes(router)
	log.Println("Server Started at localhost:8080")
	http.ListenAndServe(":8080", router)
}
