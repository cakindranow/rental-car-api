package domain

import "gorm.io/gorm"

type Product struct {
	ID        string `gorm:"primaryKey;column:id;<-:create"`
	Name      string
	Price     int64
	ImageUrl  string
	DeletedAt gorm.DeletedAt
}

func (p *Product) TableName() string {
	return "products"
}
