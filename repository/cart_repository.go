package repository

import (
	"github.com/indrawanagung/food-order-api/model/domain"
	"github.com/indrawanagung/food-order-api/model/web"
	"gorm.io/gorm"
)

type CartRepositoryInterface interface {
	FindAll(db *gorm.DB, userID string) []web.ListCartResponse
	FindByProductAndUserID(db *gorm.DB, productID string, userID string) (domain.Cart, error)
	Save(db *gorm.DB, request domain.Cart)
	Delete(db *gorm.DB, productID string, userID string)
}
