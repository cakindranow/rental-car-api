package service

import (
	"github.com/indrawanagung/food-order-api/model/domain"
	"github.com/indrawanagung/food-order-api/model/web"
)

type UserServiceInterface interface {
	Save(request web.UserCreateOrUpdateRequest) string
	Update(ID string, request web.UserCreateOrUpdateRequest)
	FindByID(ID string) web.UserResponse
	FindAll() []domain.User
}
