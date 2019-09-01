package data

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

var database *mongo.Database

func Test() {

}

func init() {
	clientOptions := options.Client().ApplyURI("mongodb://root:example@localhost:27017")
	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")

	database = client.Database("kubeui")
	fmt.Println("Connected to database:" + database.Name())

	iserted, err := ApplicationCollection().InsertOne(context.TODO(), Application{"FirstApplication"})

	if err != nil {
		panic(err.Error())
	}

	fmt.Println(iserted)

}

func ApplicationCollection() *mongo.Collection {
	return database.Collection("applications")
}

func NamespaceCollection() *mongo.Collection {
	return database.Collection("namespaces")
}

func DeploymentCollection() *mongo.Collection {
	return database.Collection("deployments")
}
