package models

import (
	"myapp/db"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name string
	Age  uint
	Work string
}

func NewUser(name, work string, age uint) {
	db.Db.Create(&User{Name: name, Work: work, Age: age})
}
