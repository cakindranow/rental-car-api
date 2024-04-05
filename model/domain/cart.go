package domain

type Cart struct {
	ProductID string `gorm:"primaryKey;column:product_id;<-:create" json:"product_id"`
	UserID    string `gorm:"primaryKey;column:user_id;<-:create" json:"user_id"`
	Total     int    `json:"total"`
}

func (c *Cart) TableName() string {
	return "carts"
}
