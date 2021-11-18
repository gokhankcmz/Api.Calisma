package ResponseModels

import (
	"time"
)

type OrderResponseDto struct {
	Status		string
	Quantity	int
	Price 		float64
	CreatedAt	time.Time
	UpdatedAt	time.Time
}