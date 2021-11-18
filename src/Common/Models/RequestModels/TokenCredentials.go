package RequestModels

type TokenCredentials struct {
	ID		string `json:"id" validate:"required,gte=2"`
	Email	string `json:"Email" validate:"required,gte=2"`
}
