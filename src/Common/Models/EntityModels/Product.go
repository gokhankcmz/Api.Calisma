package Entities

type Product struct {
	ID 			string `json:"productid,omitempty" validate:"required"`
	ImageURL    string `json:"imageurl,omitempty" validate:"required"`
	Name	    string `json:"name,omitempty" validate:"required"`
}
