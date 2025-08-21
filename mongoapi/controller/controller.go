package controller

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const connectionString = "mongodb+srv://aashishas16:Ashish%40123@cluster0.uec2coa.mongodb.net/?retryWrites=true&w=majority&appName=Cluster0" // fixed encoding
const dbName = "netflix"                                                                                                                     // Database name
const colName = "watchlist"                                                                                                                  // Collection name
const port = ":8080"                                                                                                                         // Port for the HTTP server

var Collection *mongo.Collection // Global variable to hold the collection

func init() {
	// Connect to MongoDB
	clientOptions := options.Client().ApplyURI(connectionString)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	// Ping the database to verify connection
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal("Failed to ping MongoDB:", err)
	}

	fmt.Println("MongoDB connection established successfully")

	// Set the collection
	Collection = client.Database(dbName).Collection(colName)
	fmt.Println("Collection instance is ready")
}
