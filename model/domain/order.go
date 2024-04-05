package domain

type Order struct {
	ID          string `gorm:"primaryKey;column:id;<-:create" json:"id"`
	StatusID    string `json:"status_id"`
	CarsID      string `json:"cars_id"`
	StartDate   string `json:"start_date"`
	EndDate     string `json:"end_date"`
	TotalDay    int    `json:"total_day"`
	OrderedAt   string `json:"ordered_at"`
	RequestedBy string `json:"requested_by"`
	NoteAdmin   string `json:"note_admin"`
	Cost        int64  `json:"cost"`
}

func (o *Order) TableName() string {
	return "orders"
}
