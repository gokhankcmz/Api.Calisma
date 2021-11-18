package Mapper

import (
	"Api.Calisma/src/Common/Models/EntityModels"
	"Api.Calisma/src/Common/Models/RequestModels"
)

func MapOrderDtoToOrder(OrderDTO RequestModels.CreateOrderDto) Entities.Order {
	val, _ := OrderDTO.Price.Float64()
	return Entities.Order{
		Status:     "Pending",
		Quantity:   OrderDTO.Quantity,
		Product:    OrderDTO.Product,
		Price: 		val,
	}
}
