package Entities

import "go.mongodb.org/mongo-driver/bson/primitive"

type Order struct {
	ID			primitive.ObjectID			`json:"_id,omitempty" bson:"_id,omitempty"`
	CustomerID	string			`json:"customerid,omitempty" bson:"customerid,omitempty"`
	Status		string  		`json:"status, omitempty"`
	Quantity	int				`json:"quantity, omitempty"`
	Price		float64			`json:"price" validate:"required, min=1"`
	Product		Product 		`json:"product"`
	Address		Address 		`json:"address"`
	Document					`json:"document"`
}



