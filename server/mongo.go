package server

import (
	"context"
	"log"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

/****************************************.
*	MongoDB getConnection
* 	MongoDb getCollection
* 	MongoDb  closeConnection
*  	TODO: Logging,Refactoring
******************************************/

/********
*	connect To MongoDb and get Collection by passing documentName,tableName as a String
* 	and Return the collection pointer
*********/

var client *mongo.Client

func ConnectAndGetMongoDbCollection(documentName string, tableName string) *mongo.Collection {
	// Check the connection
	collection := client.Database(documentName).Collection(tableName)
	return collection
}

func ConnectMongoDbClient() {
	clientOptions := options.Client().ApplyURI("mongodb+srv://ameyasalagre:pass12345@cluster0-prqsb.mongodb.net/")
	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)
	// Check the connection
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal("Error While Connection", err)
	}
	log.Println("Connected to MongoDB!")
}
