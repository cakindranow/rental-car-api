package web

type LoginRequest struct {
	Email    string `validate:"required,min=1,max=50" json:"email"`
	Password string `validate:"required,min=1,max=50" json:"password"`
}

type LoginResponse struct {
	Name    string `json:"name"`
	Email   string `json:"email"`
	Token   string `json:"token"`
	IsAdmin bool   `json:"is_admin"`
}
