package database

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var uri string = "mongodb+srv://yusufkarback64:9DtM5ATPlocRRVtR@cluster0.ikpoxve.mongodb.net/?retryWrites=true&w=majority"

func ConnectToDatabase() *mongo.Client {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}
	return client
}
