package domain

import "gorm.io/gorm"

type Car struct {
	ID              string `gorm:"primaryKey;column:id;<-:create"`
	Brand           string
	Model           string
	Plat            string
	DailyRentalRate int64
	ImageUrl        string
	Desc            string
	DeletedAt       gorm.DeletedAt
}

func (c *Car) TableName() string {
	return "cars"
}
