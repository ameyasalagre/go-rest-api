package utils

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

var client *mongo.Client

/********
*	connect To MongoDb and get Collection by passing documentName,tableName as a String
* 	and Return the collection pointer
*********/
func ConnectAndGetMongoDbCollection(documentName string, tableName string) *mongo.Collection {
	// get a MongoDBClient
	client := getClient()
	log.Println("Connected to MongoDB!")
	// Check the connection
	collection := client.Database(documentName).Collection(tableName)
	return collection
}

/********
*	Close  MongoDb connection
*	TODO:
*	~Stop MongoClient when its ideal for more than 2 mins
*********/
func CloseMongoDbClient() {
	client := getClient()
	client.Disconnect(context.TODO())
	log.Println("Connection to MongoDB closed.")
}


/********
*	get MongoDb client
*********/
func getClient() *mongo.Client {
	clientOptions := options.Client().ApplyURI("mongodb+srv://ameyasalagre:pass12345@cluster0-prqsb.mongodb.net/")
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
		}
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal("Error While Connection", err)
	}
	return client 
}




