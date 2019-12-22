package controllers

import (
	"fmt"
	"net/http"
)

/****************************************.
*	path : "/user"
*	purpose : List all users
*	To DO: Fetch all users from MongoDb
*
******************************************/
func User(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome User!")

}
