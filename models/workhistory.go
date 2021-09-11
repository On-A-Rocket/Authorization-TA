package models

type (
	WorkHistory struct {
		Id             string `json:"id"  validate:"required"`
		Name           string `json:"name"  validate:"required"`
		Deparment_code string `json:"deparment_code"  validate:"required"`
		Deparment_name string `json:"deparment_name" validate:"required"`
		Start_time     string `json:"start_time" validate:"required"`
		End_time       string `json:"end_time"  validate:"required"`
		Work_time      string `json:"work_time"  validate:"required"`
		Update_time    string `json:"update_time"  validate:"required"`
		Message        string `json:"message" validate:"required"`
		Work_type_code string `json:"work_type_code"  validate:"required"`
		Work_type_name string `json:"work_type_name"  validate:"required"`
		//Email string `json:"email" example:"aaa@aaa.com"  validate:"required,email"`
	}
)
