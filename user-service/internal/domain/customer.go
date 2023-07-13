package domain

type CustomerInput struct {
	FirstName string `json:"firstName" validate:"required"`
	LastName  string `json:"lastName" validate:"required"`
	Email     string `json:"email" validate:"required"`
	Phone     string `json:"phone"`
	TwitterID string `json:"twitterID"`
}
