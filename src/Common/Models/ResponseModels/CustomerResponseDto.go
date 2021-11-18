package ResponseModels

import "time"

type CustomerResponseDto struct {
	Name 		string
	Email		string
	CreatedAt	time.Time
	UpdatedAt	time.Time
	Valid		bool
}
