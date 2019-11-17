/*************************************.
*	Author : Ameya Salagre
*	github : github.com/ameyasalagre
*	Entry Point of Repository
**************************************/
package main

import (
	"fmt"
	"github.com/go-rest-api/server"
)
/************************.
*	Call to The Server
************************/
func main() {
	fmt.Print("Hello Rest Api")
	server.Start()
}
