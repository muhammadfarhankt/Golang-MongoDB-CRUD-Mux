package models

import (
	"context"
	"fmt"

	"github.com/muhammadfarhankt/Golang-MongoDB-CRUD-Mux/pkg/config"
	"go.mongodb.org/mongo-driver/bson"
)

type Intern struct {
	InternId  string   `json:"internid,omitempty" bson:"Internid,omitempty"`
	FirstName string   `json:"firstname,omitempty" bson:"Firstname,omitempty"`
	LastName  string   `json:"lastname,omitempty" bson:"Lastname,omitempty"`
	Email     string   `json:"email,omitempty" bson:"Email,omitempty"`
	Phone     string   `json:"phone,omitempty" bson:"Phone,omitempty"`
	Age       int      `json:"age,omitempty" bson:"Age,omitempty"`
	Domain    string   `json:"domain,omitempty" bson:"Domain,omitempty"`
	Address   *Address `json:"address,omitempty" bson:"Address,omitempty"`
}

type Address struct {
	City    string `json:"city,omitempty" bson:"City,omitempty"`
	Pincode string `json:"pincode,omitempty" bson:"Pincode,omitempty"`
	State   string `json:"state,omitempty" bson:"State,omitempty"`
	Country string `json:"country,omitempty" bson:"Country,omitempty"`
}

var collection = config.InitDB().Database("interns").Collection("details")

func (intern *Intern) CreateIntern() *Intern {
	// collection := ConnectionClient.Database("interns").Collection("details")
	_, err := collection.InsertOne(context.TODO(), intern)
	if err != nil {
		fmt.Println("Error while inserting the intern details: ", err)
		return nil
	}
	return intern
}

func GetAllInternDetails() []Intern {
	cursor, err := collection.Find(context.TODO(), bson.M{})
	if err != nil {
		fmt.Println("Error while fetching the intern details: ", err)
		return nil
	}
	var Interns []Intern
	if err := cursor.All(context.TODO(), &Interns); err != nil {
		fmt.Println("Error while decoding the intern details: ", err)
		return nil
	}
	//fmt.Println("Interns: ", Interns)
	return Interns
}

func GetInternDetails(id string) *Intern {
	//fmt.Println("Get intern details")
	var intern Intern
	// fmt.Println("id: ", id)
	cursor := collection.FindOne(context.TODO(), bson.M{"Internid": id})
	cursor.Decode(&intern)
	return &intern
}

func DeleteIntern(id string) *Intern {
	//fmt.Println("Delete intern")
	var intern Intern
	// fmt.Println("id: ", id)
	cursor := collection.FindOneAndDelete(context.TODO(), bson.M{"Internid": id})
	cursor.Decode(&intern)
	return &intern
}

func (i *Intern) UpdateInternDetails(id string) *Intern {
	// fmt.Println("Update intern")
	// fmt.Println("id: ", id)
	var intern Intern
	update := bson.M{
		"$set": i,
	}
	//fmt.Println("update: ", update)
	_, err := collection.UpdateOne(context.TODO(), bson.M{"Internid": id}, update)
	//fmt.Println("cursor: ", cursor)
	if err != nil {
		fmt.Println("Error while updating the intern details: ", err)
		return nil
	}
	if err := collection.FindOne(context.TODO(), bson.M{"Internid": id}).Decode(&intern); err != nil {
		fmt.Println("Error while decoding the intern details: ", err)
		return nil
	}
	//fmt.Println("Intern: ", intern)
	return &intern
}
