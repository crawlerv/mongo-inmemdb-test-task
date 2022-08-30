package dto

type CreateInput struct {
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Age         uint8  `json:"age"`
	CardNumber  uint32 `json:"card_number"`
	PhoneNumber string `json:"phone_number"`
	Verified    bool   `json:"verified"`
}

type UpdateInput struct {
	FirstName   *string `json:"first_name"`
	LastName    *string `json:"last_name"`
	Age         *uint8  `json:"age"`
	CardNumber  *uint32 `json:"card_number"`
	PhoneNumber *string `json:"phone_number"`
	Verified    *bool   `json:"verified"`
}
