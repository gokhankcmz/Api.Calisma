package Mapper

import (
	Entities "Api.Calisma/src/Common/Models/EntityModels"
	"Api.Calisma/src/Common/Models/ResponseModels"
)

func GetCustomerResponse(Customer Entities.Customer) ResponseModels.CustomerResponseDto {
	return ResponseModels.CustomerResponseDto{
		Name:      Customer.Name,
		Email:     Customer.Email,
		CreatedAt: Customer.CreatedAt,
		UpdatedAt: Customer.UpdatedAt,
		Valid:     Customer.Valid,
	}
}
