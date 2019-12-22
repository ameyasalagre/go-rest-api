package controllers

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/go-rest-api/modules"
	"github.com/go-rest-api/utils"
	"go.mongodb.org/mongo-driver/bson"
)

/****************************************.
*	path : "/"
*	purpose : Test Page
*	TO DO : ------
*
******************************************/
// You will be using this Trainer type later in the program
type Trainer struct {
	ip       string
	Username string
	Password string
}

func Home(w http.ResponseWriter, r *http.Request) {
	collection :=  utils.ConnectAndGetMongoDbCollection("LoginApp", "users")
	filter := bson.M{"ip": "192.168.0.103"}
	result := modules.User{}
	err := collection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		log.Fatal(err)
	}
	go utils.CloseMongoDbClient()
	fmt.Fprintf(w, result.Username)

}
