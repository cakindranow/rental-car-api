package repository

import (
	"github.com/indrawanagung/food-order-api/model/domain"
	"gorm.io/gorm"
)

type ProductRepositoryInterface interface {
	Save(db *gorm.DB, product domain.Product)
	FindAll(db *gorm.DB) []domain.Product
	FindById(db *gorm.DB, productId string) (domain.Product, error)
}
