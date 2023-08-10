package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Employee struct represents the employee data
type Employee struct {
	ID   string `bson:"_id,omitempty"`
	Name string `bson:"name"`
	Age  int    `bson:"age"`
	//CreatedAt time.Time `bson:"createdAt,omitempty"`
	//UpdatedAt time.Time `bson:"updatedAt,omitempty"`
}

// EmployeeRepository is the interface for employee-related database operations
type EmployeeRepository interface {
	SaveEmployee(employee *Employee) error
	GetEmployeeByID(id string) (*Employee, error)
}

// MongoDBEmployeeRepository implements the EmployeeRepository interface using MongoDB
type MongoDBEmployeeRepository struct {
	//type Mydata struct {
	collection *mongo.Collection
}

// SaveEmployee saves an employee to the database
func (r *MongoDBEmployeeRepository) SaveEmployee(employee *Employee) error {
	//employee.CreatedAt = time.Now()
	//employee.UpdatedAt = time.Now()

	_, err := r.collection.InsertOne(context.Background(), employee)
	if err != nil {
		return err
	}

	return nil
}

// GetEmployeeByID retrieves an employee from the database by ID
func (r *MongoDBEmployeeRepository) GetEmployeeByID(id string) (*Employee, error) {
	var employee Employee

	err := r.collection.FindOne(context.Background(), bson.M{"_id": id}).Decode(&employee)
	if err != nil {
		return nil, err
	}

	return &employee, nil
}

func main() {
	// Set up the MongoDB client
	clientOptions := options.Client().ApplyURI("mongodb+srv://Ganesh:12345@mydbcluster.xwabpqc.mongodb.net/?retryWrites=true&w=majority")
	client, err := mongo.NewClient(clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	// Connect to the MongoDB cluster
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	// Select the MongoDB database and collection
	db := client.Database("Mydatabase1")
	collection := db.Collection("emp1")

	// Create an instance of the MongoDBEmployeeRepository
	repo := &MongoDBEmployeeRepository{
		collection: collection,
	}

	// Example usage: save an employee
	employee := &Employee{
		Name: "Ajay",
		Age:  26,
	}
	err = repo.SaveEmployee(employee)
	if err != nil {
		log.Fatal(err)
	}

	// Example usage: get an employee
	filter := bson.M{"age": bson.M{"$gt": 21}}
	cur, err := collection.Find(ctx, filter)
	if err != nil {
		log.Fatal(err)
	}
	defer cur.Close(ctx)
	for cur.Next(ctx) {
		//var employee Employee
		err := cur.Decode(&employee)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Fetched Employee:\n", employee.ID, employee.Name)
	}
}
