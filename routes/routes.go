package routes

import (
	"github.com/go-rest-api/controllers"
	"github.com/gorilla/mux"
)

//Add New App Routes in below function
func IntiateRoutes(routes *mux.Router) {
	routes.HandleFunc("/", controllers.Home)
	routes.HandleFunc("/user", controllers.User)
}
