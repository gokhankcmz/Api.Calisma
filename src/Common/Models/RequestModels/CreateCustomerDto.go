package RequestModels

import (
	"Api.Calisma/src/Common/Models/EntityModels"
)

type CreateCustomerDto struct {
	Name    string            `json:"name" validate:"required,gte=2"`
	Address Entities.Address `json:"address" validate:"required"`
	Email   string            `json:"email" validate:"required,email,gte=6"`
}
