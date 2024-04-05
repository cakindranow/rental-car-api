package web

type CreateOrUpdateProduct struct {
	Name     string `validate:"required" json:"name" form:"name"`
	Price    int64  `validate:"required" json:"price" form:"price"`
	ImageUrl string `validate:"required" json:"image_url" form:"image_url"`
}

type ProductResponse struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	ImageUrl string `json:"image_url"`
	Price    int64  `json:"price"`
}
