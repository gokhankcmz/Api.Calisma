package Mapper

import (
	Entities "Api.Calisma/src/Common/Models/EntityModels"
	"Api.Calisma/src/Common/Models/ResponseModels"
)

func MapOrderToOrderResponseDto(OrderEntity *Entities.Order) ResponseModels.OrderResponseDto {
	return ResponseModels.OrderResponseDto{
		Status:    OrderEntity.Status,
		Quantity:  OrderEntity.Quantity,
		Price:     OrderEntity.Price,
		CreatedAt: OrderEntity.CreatedAt,
		UpdatedAt: OrderEntity.UpdatedAt,
	}
}

