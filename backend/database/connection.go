package database

import (
	"context"
	"fmt"

	// "go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var DB *mongo.Client
var UserCollection *mongo.Collection
var ServicesCollection *mongo.Collection
var ServiceDayDetailsCollection *mongo.Collection

func Connect(uri string)  {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}	
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		panic(err)
	}
	fmt.Println("Connected to Database!")
	UserCollection= client.Database("appoiment-db").Collection("user")
	ServicesCollection= client.Database("appoiment-db").Collection("services")
	ServiceDayDetailsCollection= client.Database("appoiment-db").Collection("service_day_details")
	DB=client
}


func Disconnect(){

	if err := DB.Disconnect(context.TODO()); err != nil {
		panic(err)
	}
	fmt.Println("Database DisConnected!")
}

