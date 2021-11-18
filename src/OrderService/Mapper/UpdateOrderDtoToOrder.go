package Mapper

import (
	"Api.Calisma/src/Common/Models/EntityModels"
	"Api.Calisma/src/Common/Models/RequestModels"
)

func ApplyChangesToOrder(UpdateDTO *RequestModels.UpdateOrderDto, orderEntity *Entities.Order){
	orderEntity.Product = UpdateDTO.Product
	orderEntity.Quantity = UpdateDTO.Quantity
	val, _ := UpdateDTO.Price.Float64()
	orderEntity.Price = val
}