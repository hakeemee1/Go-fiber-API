package models


type Register struct {
	Email    string `json:"email" validate:"required,email"`
	Username string `json:"username" validate:"required,username_validate,alphanum,min=1,max=30"`
	Password string `json:"password" validate:"required,min=6,max=20"`
	LineId   string `json:"lineId" validate:"required"`
	Phone    string `json:"phone" validate:"required,numeric,len=10"`
	Business string `json:"business" validate:"required"`
	WebName  string `json:"webName" validate:"required,min=2,max=30,web_validate"`
}
