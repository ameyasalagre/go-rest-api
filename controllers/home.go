package controllers

import(
	"net/http"
	"fmt"
)

/****************************************.
*	path : "/"
*	purpose : Test Page
*	TO DO : ------
*
******************************************/
func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome home!")
}