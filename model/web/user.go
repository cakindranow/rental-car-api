package web

import "github.com/indrawanagung/food-order-api/model/domain"

type UserCreateOrUpdateRequest struct {
	Name     string `validate:"required,min=5,max=50" json:"name"`
	Email    string `validate:"required,min=5,max=50" json:"email"`
	Password string `validate:"required,min=5,max=50" json:"password"`
	Phone    string `validate:"required,min=5,max=50" json:"phone"`
	Address  string `validate:"required,min=5,max=50" json:"address"`
	SIM      string `validate:"required,min=5,max=50" json:"sim"`
}

type UserResponse struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Email   string `json:"email"`
	Phone   string `json:"phone"`
	Address string `json:"address"`
	SIM     string `json:"sim"`
}

func ToUserResponse(user domain.User) UserResponse {
	return UserResponse{
		ID:      user.ID,
		Name:    user.Name,
		Email:   user.Email,
		Phone:   user.Phone,
		Address: user.Address,
		SIM:     user.SIM,
	}
}
