package Mapper

import (
	"Api.Calisma/src/Common/Models/EntityModels"
	"Api.Calisma/src/Common/Models/RequestModels"
)

func UpdateCustomerDtoToCustomer(CustomerDTO RequestModels.UpdateCustomerDto) Entities.Customer {
	return Entities.Customer{
		Document: Entities.Document{},
		Address:  CustomerDTO.Address,
		Name:     CustomerDTO.Name,
		Email:    CustomerDTO.Email,
		Valid:    false,
	}
}

func AppylChangesToCustomer(UpdateDTO RequestModels.UpdateCustomerDto, customerEntity *Entities.Customer){

	if UpdateDTO.Address.AddressLine != ""{
		customerEntity.Address.AddressLine = UpdateDTO.Address.AddressLine
	}
	if UpdateDTO.Address.City != ""{
		customerEntity.Address.City = UpdateDTO.Address.City
	}
	if UpdateDTO.Address.Country!= ""{
		customerEntity.Address.Country = UpdateDTO.Address.Country
	}
	if UpdateDTO.Address.CityCode != 0{
		customerEntity.Address.CityCode = UpdateDTO.Address.CityCode
	}
	if UpdateDTO.Name != ""{
		customerEntity.Name = UpdateDTO.Name
	}
	if UpdateDTO.Email != ""{
		customerEntity.Email = UpdateDTO.Email
	}
}