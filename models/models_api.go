package models

import "gorm.io/gorm"

type Register struct {
	Email    string `json:"email" validate:"required,email"`
	Username string `json:"username" validate:"required,username_validate,alphanum,min=1,max=30"`
	Password string `json:"password" validate:"required,min=6,max=20"`
	LineId   string `json:"lineId" validate:"required"`
	Phone    string `json:"phone" validate:"required,numeric,len=10"`
	Business string `json:"business" validate:"required"`
	WebName  string `json:"webName" validate:"required,min=2,max=30,web_validate"`
}

// gorm.Model เป็นคลาสที่ใช้ในการให้ฟิลด์เพิ่มเติมในตาราง เช่น  id, created_at, updated_at, deleted_at
// ฟิลด์เหล่านี้จะถูกสร้างขึ้นมาอัตโนมัติโดย gorm เมื่อเราทำการ Insert, Update, Delete ในตาราง
type Dogs struct {
	gorm.Model
	Name  string `json:"name"`
	DogID int    `json:"dog_id"`
}

type DogsRes struct {
	Name  string `json:"name"`
	DogID int    `json:"dog_id"`
	Type  string `json:"type"`
}

type Companies struct {
	gorm.Model
	Name string `json:"name"`
	Web  string `json:"web"`
	Line string `json:"line"`
}

type ResultData struct {
	Data        []DogsRes `json:"data"`
	Name        string    `json:"name"`
	// Count       int       `json:"count"`
	Sum_red     int       `json:"sum_red"`
	Sum_green   int       `json:"sum_green"`
	Sum_pink    int       `json:"sum_pink"`
	Sum_noColor int       `json:"sum_nocolor"`
}

//User Profile
type UserProfiles struct {
	gorm.Model
	EmployeeID string `gorm:"unique" json:"employee_id"`
	Name       string `json:"name" validate:"required"`
	LastName   string `json:"lastname" validate:"required"`
	Birthday   string `json:"birthday" validate:"required"`
	Age        int    `json:"age" validate:"required,min=18"`
	Email      string `json:"email" validate:"required,email"`
	Tel        string `json:"tel" validate:"required"`
}