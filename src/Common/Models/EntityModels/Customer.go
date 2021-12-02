package Entities

import "go.mongodb.org/mongo-driver/bson/primitive"

type Customer struct {
	Document `json:"document" validate:"required"`
	Address  Address            `json:"address"`
	ID       primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty" validate:"required"`
	Name     string             `json:"name,omitempty" validate:"required"`
	Email    string             `json:"email,omitempty" validate:"required"`
	Valid    bool               `json:"valid,omitempty" validate:"required"`
}
