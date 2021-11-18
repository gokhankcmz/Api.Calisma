package ResponseModels

type AddressDto struct{
	AddressLine string `json:"addressline,omitempty" validate:"required"`
	City        string `json:"city,omitempty" validate:"required"`
	Country     string `json:"country,omitempty" validate:"required"`
	CityCode    int    `json:"citycode,omitempty" validate:"required"`
}
