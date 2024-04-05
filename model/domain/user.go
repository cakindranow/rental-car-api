package domain

import "gorm.io/gorm"

type User struct {
	ID        string `gorm:"primary_key;column:id" json:"id"`
	Name      string `gorm:"column:name" json:"name"`
	Email     string `gorm:"column:email" json:"email"`
	Password  string `gorm:"column:password" json:"password"`
	Phone     string `gorm:"column:phone" json:"phone"`
	Address   string `gorm:"column:address" json:"address"`
	SIM       string `gorm:"column:sim" json:"sim"`
	IsAdmin   bool   `gorm:"column:is_admin" json:"is_admin"`
	DeletedAt gorm.DeletedAt
}

func (u *User) TableName() string {
	return "users"
}
