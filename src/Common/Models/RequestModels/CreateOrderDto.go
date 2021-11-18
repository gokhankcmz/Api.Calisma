package RequestModels

import (
	Entities "Api.Calisma/src/Common/Models/EntityModels"
	"github.com/shopspring/decimal"
)

type CreateOrderDto struct {
	Quantity	int					`json:"quantity" validate:"required,min=1"`
	Price		decimal.Decimal		`json:"price" validate:"required,min=1"`
	Product		Entities.Product	`json:"product" validate:"required"`
}

