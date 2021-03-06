package models

type (
	User struct {
		Name  string `json:"name"  validate:"required"`
		Email string `json:"email" validate:"required,email"`
	}
)
