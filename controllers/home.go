package controllers

import(
	"net/http"
	"fmt"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome home!")
}