package service

import (
	"github.com/indrawanagung/food-order-api/model/domain"
	"github.com/indrawanagung/food-order-api/model/web"
)

type CartServiceInterface interface {
	FindAll(userID string) []web.ListCartResponse
	FindByProductAndUserID(productID string, userID string) domain.Cart
	Save(request web.CartCreateRequest)
	Delete(productID string, userID string)
}
