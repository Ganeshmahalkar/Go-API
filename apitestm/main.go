package main

import (
	"apitestm/dbiface"
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Employee struct represents the employee data
type User struct {
	Name string `bson:"name"`
	Age  int    `bson:"age"`
}

func insertData(collection dbiface.CollectionAPI, user User) (*mongo.InsertOneResult, error) {
	res, err := collection.InsertOne(context.Background(), user)
	if err != nil {
		return res, err
	}
	return res, nil
}

func deleteData(collection dbiface.CollectionAPI, user User) (*mongo.DeleteResult, error) {
	res, err := collection.DeleteOne(context.Background(), user)
	if err != nil {
		return res, err
	}
	return res, nil
}

func main() {
	c, err := mongo.Connect(context.Background(), options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatal("Error : %", err)
	}
	db := c.Database("Employee")
	col := db.Collection("details")
	//col.DeleteOne
	res, err := insertData(col, User{"Rajesh", 32})
	log.Println(res, err)

}
