package Entities

import "go.mongodb.org/mongo-driver/bson/primitive"

type Customer struct {
	Document `json:"document"`
	Address  Address `json:"address"`
	ID       primitive.ObjectID  `json:"_id,omitempty" bson:"_id,omitempty"`
	Name     string  `json:"name,omitempty"`
	Email    string  `json:"email,omitempty"`
	Valid    bool    `json:"valid,omitempty"`
}


