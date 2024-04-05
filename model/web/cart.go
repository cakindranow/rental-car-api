package web

type CartCreateRequest struct {
	Total     int    `validate:"gt=0" json:"total"`
	ProductID string `validate:"required,min=1,max=50" json:"product_id"`
	UserID    string `validate:"required,min=1,max=50" json:"user_id"`
}

type ListCartResponse struct {
	UserID      string `json:"user_id"`
	ProductID   string `json:"product_id"`
	Total       int    `json:"total"`
	ProductName string `json:"product_name"`
	Price       int64  `json:"price"`
	ImageUrl    string `json:"image_url"`
}
