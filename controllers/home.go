package controllers

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"github.com/go-rest-api/modules"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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
	// fmt.Fprintf(w, "Welcome home!")
	// Set client options
	clientOptions := options.Client().ApplyURI("mongodb+srv://ameyasalagre:pass12345@cluster0-prqsb.mongodb.net/")

	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)

	collection := client.Database("LoginApp").Collection("users")

	filter := bson.M{"ip": "192.168.0.103"}
	result := modules.User{}

	err = collection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Found a single document: %+v\n", result)

	fmt.Fprintf(w, result.Password)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")
}
