package main

import (
	// "context"
	"context"
	"dashboard_pnj/database"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Dosen struct {
	Nama string
	Umur int
}

func CreateCollection(client *mongo.Client) {
	createOpts := options.CreateCollection().SetCapped(false) // Change options as needed
	err := client.Database("dashboard_Pnj").CreateCollection(context.TODO(), "akademi", createOpts)
	if err != nil {
		panic(err)
	}
	fmt.Println("Collection created successfully.")

}

// 9DtM5ATPlocRRVtR
func main() {
	client := database.ConnectToDatabase()

	// //insert data to collection
	// coll := client.Database("dashboard_Pnj").Collection("dosen")
	// newDosen := Dosen{Nama: "agung", Umur: 22}
	// _, err := coll.InsertOne(context.TODO(), newDosen)
	// if err != nil {
	// 	panic(err)
	// }
	CreateCollection(client)
}
