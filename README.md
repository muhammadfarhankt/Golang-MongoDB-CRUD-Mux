# Building a RESTful CRUD API in Golang with MongoDB and Gorilla/Mux
In this tutorial, we will build a Golang CRUD (Create, Read, Update, Delete) application using MongoDB as our database and Gorilla Mux for routing. This application will manage interns' details, allowing us to perform basic CRUD operations.

This tutorial provides a comprehensive guide to building a CRUD API using Golang, MongoDB, and Gorilla Mux. It covers setting up the project structure, defining models, implementing CRUD operations, defining routes, and starting the server. Each step is accompanied by detailed comments to explain the purpose and functionality of the code.

### Prerequisites
- Golang
- MongoDB
  
### Project Structure
Before diving into the code, let's take a look at the structure of our project:

```bash
Golang-MongoDB-CRUD-Mux/
├── cmd/
│   └── main/
│       └── main.go
│       └── .env
├── pkg/
│   ├── config/
│   │   └── config.go
│   ├── controller/
│   │   └── controller.go
│   ├── models/
│   │   └── intern.go
│   ├── routes/
│   │   └── route.go
│   └── utils/
        └── util.go
```

### Start using it

```bash
go mod init <project_name>
```

Download and install dependencies:

```bash
go get github.com/gorilla/mux
```

```bash
go get github.com/joho/godotenv
```

```bash
go get go.mongodb.org/mongo-driver/mongo
```

## 1. Setting Up MongoDB Connection
We'll start by configuring our MongoDB connection in the config.go file. This file will handle the initialization of our MongoDB client.

```bash
// package/config/config.go

package config

import (
    "context"
    "log"
    "os"

    "github.com/joho/godotenv"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
)

func InitDB() *mongo.Client {
    // Load environment variables
    if err := godotenv.Load(); err != nil {
        log.Println("No .env file found")
    }

    // Get the MongoDB URI
    dbURI := os.Getenv("MONGODB_URI")
    if dbURI == "" {
        log.Fatal("MONGODB_URI is not set")
    }

    // Set client options
    clientOptions := options.Client().ApplyURI(dbURI)

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

    log.Println("Connected to MongoDB!")
    return client
}

```

## 2. Defining the Intern Model
Next, let's define our Intern model in the intern.go file under the models package.
```bash
// package/models/intern.go

package models

import (
    "context"
    "fmt"

    "github.com/muhammadfarhankt/Golang-MongoDB-CRUD-Mux/pkg/config"
    "go.mongodb.org/mongo-driver/bson"
)

// Intern represents the structure of an intern
type Intern struct {
    InternID  string   `json:"internid,omitempty" bson:"Internid,omitempty"`
    FirstName string   `json:"firstname,omitempty" bson:"Firstname,omitempty"`
    LastName  string   `json:"lastname,omitempty" bson:"Lastname,omitempty"`
    Email     string   `json:"email,omitempty" bson:"Email,omitempty"`
    Phone     string   `json:"phone,omitempty" bson:"Phone,omitempty"`
    Age       int      `json:"age,omitempty" bson:"Age,omitempty"`
    Domain    string   `json:"domain,omitempty" bson:"Domain,omitempty"`
    Address   *Address `json:"address,omitempty" bson:"Address,omitempty"`
}

// Address represents the address of an intern
type Address struct {
    City    string `json:"city,omitempty" bson:"City,omitempty"`
    Pincode string `json:"pincode,omitempty" bson:"Pincode,omitempty"`
    State   string `json:"state,omitempty" bson:"State,omitempty"`
    Country string `json:"country,omitempty" bson:"Country,omitempty"`
}

// MongoDB collection instance
var collection = config.InitDB().Database("interns").Collection("details")

// CRUD operations...

```

##  3. Implementing CRUD Operations
Now, let's implement the CRUD operations in the controller.go file under the controller package:
```bash
// package/controller/controller.go

package controller

import (
    "encoding/json"
    "net/http"

    "github.com/gorilla/mux"
    "github.com/muhammadfarhankt/Golang-MongoDB-CRUD-Mux/pkg/models"
    "github.com/muhammadfarhankt/Golang-MongoDB-CRUD-Mux/pkg/utils"
)

// GetAllInternDetails handles the HTTP GET request to fetch all intern details
func GetAllInternDetails(w http.ResponseWriter, r *http.Request) {
    // Fetch all intern details
    allInterDetails := models.GetAllInternDetails()

    // Marshal intern details into JSON
    res, _ := json.Marshal(allInterDetails)

    // Set response header
    w.Header().Set("Content-Type", "application/json")

    // Write response
    if allInterDetails == nil {
        w.WriteHeader(http.StatusNotFound)
        res = []byte(`{"message": "Interns not found"}`)
    } else {
        w.WriteHeader(http.StatusOK)
    }
    w.Write(res)
}

// GetInternDetails handles the HTTP GET request to fetch details of a specific intern by their ID
func GetInternDetails(w http.ResponseWriter, r *http.Request) {
    // Parse intern ID from request URL
    vars := mux.Vars(r)
    id := vars["id"]

    // Fetch intern details by ID
    InternDetails := models.GetInternDetails(id)

    // Marshal intern details into JSON
    res, _ := json.Marshal(InternDetails)

    // Set response header
    w.Header().Set("Content-Type", "application/json")

    // Write response
    if InternDetails.InternId == "" {
        w.WriteHeader(http.StatusNotFound)
        res = []byte(`{"message": "Intern not found"}`)
    } else {
        w.WriteHeader(http.StatusOK)
    }
    w.Write(res)
}

// CreateIntern handles the HTTP POST request to create a new intern
func CreateIntern(w http.ResponseWriter, r *http.Request) {
    // Parse request body and decode into Intern struct
    CreateIntern := &models.Intern{}
    utils.ParseBody(r, CreateIntern)

    // Create intern
    intern := CreateIntern.CreateIntern()

    // Marshal created intern into JSON
    res, _ := json.Marshal(intern)

    // Set response header
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)

    // Write response
    w.Write(res)
}

// UpdateIntern handles the HTTP PUT request to update the details of an existing intern
func UpdateIntern(w http.ResponseWriter, r *http.Request) {
    // Parse intern ID from request URL
    vars := mux.Vars(r)
    id := vars["id"]

    // Parse request body and decode into Intern struct
    UpdateIntern := &models.Intern{}
    utils.ParseBody(r, UpdateIntern)

    // Update intern details
    internDetailsToUpdate := UpdateIntern.UpdateInternDetails(id)

    // Marshal updated intern details into JSON
    res, _ := json.Marshal(internDetailsToUpdate)

    // Set response header
    w.Header().Set("Content-Type", "application/json")

    // Write response
    if internDetailsToUpdate == nil {
        w.WriteHeader(http.StatusNotFound)
        res = []byte(`{"message": "Intern not found"}`)
    } else {
        w.WriteHeader(http.StatusOK)
    }
    w.Write(res)
}

// DeleteIntern handles the HTTP DELETE request to delete an existing intern
func DeleteIntern(w http.ResponseWriter, r *http.Request) {
    // Parse intern ID from request URL
    vars := mux.Vars(r)
    id := vars["id"]

    // Delete intern
    delete := models.DeleteIntern(id)

    // Marshal delete result into JSON
    res, _ := json.Marshal(delete)

    // Set response header
    w.Header().Set("Content-Type", "application/json")

    // Write response
    if delete.InternId == "" {
        w.WriteHeader(http.StatusNotFound)
        res = []byte(`{"message": "Intern not found"}`)
    } else {
        w.WriteHeader(http.StatusOK)
    }
    w.Write(res)
}

```

## 4. Handling Routes with Gorilla Mux
We'll use Gorilla Mux to handle our routes. Define the routes in the route.go file under the routes package.
```bash
// package/routes/route.go

package routes

import (
    "github.com/gorilla/mux"
    "github.com/muhammadfarhankt/Golang-MongoDB-CRUD-Mux/pkg/controller"
)

// InternRoutes defines routes for intern-related CRUD operations
var InternRoutes = func(r *mux.Router) {
    r.HandleFunc("/interns", controller.GetAllInternDetails).Methods("GET")
    r.HandleFunc("/intern/{id}", controller.GetInternDetails).Methods("GET")
    r.HandleFunc("/intern", controller.CreateIntern).Methods("POST")
    r.HandleFunc("/intern/{id}", controller.UpdateIntern).Methods("PUT")
    r.HandleFunc("/intern/{id}", controller.DeleteIntern).Methods("DELETE")
}
```
## 5. Handling Databasae Operations 

```bash
// package/models/intern.go

// CRUD operations...


// CreateIntern inserts a new intern into the database
func (intern *Intern) CreateIntern() *Intern {
	// Insert the intern details into the MongoDB collection
	_, err := collection.InsertOne(context.TODO(), intern)
	if err != nil {
		// If an error occurs while inserting, log the error and return nil
		fmt.Println("Error while inserting the intern details: ", err)
		return nil
	}
	// Return the created intern
	return intern
}

// GetAllInternDetails retrieves all intern details from the database
func GetAllInternDetails() []Intern {
	// Find all intern documents in the MongoDB collection
	cursor, err := collection.Find(context.TODO(), bson.M{})
	if err != nil {
		// If an error occurs while fetching, log the error and return nil
		fmt.Println("Error while fetching the intern details: ", err)
		return nil
	}
	// Decode the documents into an array of Intern structs
	var Interns []Intern
	if err := cursor.All(context.TODO(), &Interns); err != nil {
		// If an error occurs while decoding, log the error and return nil
		fmt.Println("Error while decoding the intern details: ", err)
		return nil
	}
	// Return the array of intern details
	return Interns
}

// GetInternDetails retrieves details of a specific intern by their ID from the database
func GetInternDetails(id string) *Intern {
	// Create a new Intern instance
	var intern Intern

	// Find intern details by ID in the MongoDB collection
	cursor := collection.FindOne(context.TODO(), bson.M{"Internid": id})

	// Decode the document into the intern struct
	cursor.Decode(&intern)

	// Return the intern details
	return &intern
}

// DeleteIntern deletes an existing intern from the database
func DeleteIntern(id string) *Intern {
	// Create a new Intern instance
	var intern Intern

	// Find and delete the intern by ID from the MongoDB collection
	cursor := collection.FindOneAndDelete(context.TODO(), bson.M{"Internid": id})

	// Decode the document into the intern struct
	cursor.Decode(&intern)

	// Return the deleted intern details
	return &intern
}

// UpdateInternDetails updates the details of an existing intern in the database
func (i *Intern) UpdateInternDetails(id string) *Intern {
	// Create a new Intern instance
	var intern Intern

	// Define the update query
	update := bson.M{
		"$set": i,
	}

	// Update the intern details in the MongoDB collection
	_, err := collection.UpdateOne(context.TODO(), bson.M{"Internid": id}, update)
	if err != nil {
		// If an error occurs while updating, log the error and return nil
		fmt.Println("Error while updating the intern details: ", err)
		return nil
	}

	// Fetch the updated intern details from the MongoDB collection
	if err := collection.FindOne(context.TODO(), bson.M{"Internid": id}).Decode(&intern); err != nil {
		// If an error occurs while decoding, log the error and return nil
		fmt.Println("Error while decoding the intern details: ", err)
		return nil
	}

	// Return the updated intern details
	return &intern
}
```

## 6. Main Function to Start the Server
Finally, we'll bootstrap our application in the main.go file under the main package:
```bash
// CMD/main/main.go

package main

import (
    "fmt"
    "log"
    "net/http"

    "github.com/gorilla/mux"
    "github.com/muhammadfarhankt/Golang-MongoDB-CRUD-Mux/pkg/routes"
)

func main() {
    // Create new router
    r := mux.NewRouter()

    // Define routes
    routes.InternRoutes(r)

    // Set root handler
    http.Handle("/", r)

    // Start server
    fmt.Println("Server is running on port 8080...")
    log.Fatal(http.ListenAndServe(":8080", nil))
}

```

### Conclusion
In this tutorial, we've built a Golang CRUD application using MongoDB and Gorilla Mux. We've covered setting up MongoDB connection, defining models, implementing CRUD operations, handling routes with Gorilla Mux, and bootstrapping the application. You can further expand this application by adding authentication, validation, and error handling mechanisms as per your requirements.

### Run Project

```bash
go run main.go
```
