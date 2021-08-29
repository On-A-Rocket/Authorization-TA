package model

type (
	User struct {
		Name  string `json:"name" example:"jhpark" validate:"required"`
		Email string `json:"email" example:"aaa@aaa.com"  validate:"required,email"`
	}
)
