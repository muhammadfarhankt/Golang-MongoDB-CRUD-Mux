package models

type Intern struct {
	ID        string `json:"id,omitempty" bson:"_id,omitempty"`
	FirstName string `json:"firstname,omitempty" bson:"firstname,omitempty"`
	LastName  string `json:"lastname,omitempty" bson:"lastname,omitempty"`
	Email     string `json:"email,omitempty" bson:"email,omitempty"`
	Phone     string `json:"phone,omitempty" bson:"phone,omitempty"`
	Age       int    `json:"age,omitempty" bson:"age,omitempty"`
	Domain    string `json:"domain,omitempty" bson:"domain,omitempty"`
}
