package controllers

import (
	"fmt"
	"net/http"
)

func User(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome User!")
}
