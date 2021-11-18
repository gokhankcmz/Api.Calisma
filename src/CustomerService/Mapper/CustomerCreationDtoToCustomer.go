package Mapper

import (
	"Api.Calisma/src/Common/Models/EntityModels"
	"Api.Calisma/src/Common/Models/RequestModels"
)

func CreateCustomerDtoToCustomer(CustomerDTO RequestModels.CreateCustomerDto) Entities.Customer {
	return Entities.Customer{
		Address:  CustomerDTO.Address,
		Name:     CustomerDTO.Name,
		Email:    CustomerDTO.Email,
		Valid:    false,

	}
}
