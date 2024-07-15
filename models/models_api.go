package models

import (
	"regexp"

	"github.com/go-playground/validator"
)

type Register struct {
	Email        string `json:"email" validate:"required,email"`
	Username     string `json:"username" validate:"required,alphanum,min=1,max=30"`
	Password     string `json:"password" validate:"required,min=6,max=20"`
	LineID       string `json:"line_id" validate:"omitempty"`
	PhoneNumber  string `json:"phone_number" validate:"required,numeric,len=10"`
	BusinessType string `json:"business_type" validate:"required"`
	Website      string `json:"website" validate:"required,min=2,max=30,websiteDomain"`
}

// CustomValidator เป็นฟังก์ชันสำหรับ validate ชื่อเว็บไซต์
func WebsiteDomainValidator(fl validator.FieldLevel) bool {
	website := fl.Field().String()
	regex := regexp.MustCompile(`^[a-zA-Z0-9][a-zA-Z0-9-]{0,28}[a-zA-Z0-9]$`)
	return regex.MatchString(website)
}
