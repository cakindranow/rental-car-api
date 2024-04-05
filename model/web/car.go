package web

type CreateOrUpdateCar struct {
	Brand           string `validate:"required" json:"brand" form:"brand"`
	Model           string `validate:"required" json:"model" form:"model"`
	Plat            string `validate:"required" json:"plat" form:"plat"`
	DailyRentalRate int64  `validate:"required" json:"daily_rental_rate" form:"daily_rental_rate"`
	Desc            string `validate:"required" json:"desc" form:"desc"`
	ImageUrl        string `validate:"required" json:"image_url" form:"image_url"`
}

type CreateOrderRequest struct {
	CarsID    string `validate:"required" json:"cars_id"`
	StartDate string `validate:"required" json:"start_date"`
	EndDate   string `validate:"required" json:"end_date"`
	TotalDay  int    `validate:"required" json:"total_day"`
}

type CarResponse struct {
	Id              string `json:"id"`
	Brand           string `json:"brand"`
	Model           string `json:"model" `
	Plat            string `json:"plat" `
	DailyRentalRate int64  `json:"daily_rental_rate"`
	ImageUrl        string `json:"image_url"`
	Desc            string `json:"desc"`
}

type OrderResponse struct {
	ID            string `json:"id"`
	StatusID      string `json:"status_id"`
	Status        string `json:"status_name"`
	RequestedName string `json:"requested_name"`
	CarsID        string `json:"cars_id"`
	Brand         string `json:"brand"`
	Model         string `json:"model"`
	Plat          string `json:"plat"`
	StartDate     string `json:"start_date"`
	EndDate       string `json:"end_date"`
	TotalDay      int    `json:"total_day"`
	OrderedAt     string `json:"ordered_at"`
	RequestedBy   string `json:"requested_by"`
	NoteAdmin     string `json:"note_admin"`
	Cost          int64  `json:"cost"`
}
