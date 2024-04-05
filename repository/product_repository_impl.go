package repository

import (
	"github.com/gofiber/fiber/v2/log"
	"github.com/indrawanagung/food-order-api/model/domain"
	"gorm.io/gorm"
)

type ProductRepositoryImpl struct {
}

func NewProductRepository() ProductRepositoryInterface {
	return &ProductRepositoryImpl{}
}

func (p ProductRepositoryImpl) Save(db *gorm.DB, product domain.Product) {
	err := db.Save(&product).Error
	if err != nil {
		log.Error(err)
		panic(err)
	}
}

func (p ProductRepositoryImpl) FindAll(db *gorm.DB) []domain.Product {
	var products []domain.Product
	err := db.Find(&products).Error
	if err != nil {
		log.Error(err)
		panic(err)
	}
	return products
}

func (p ProductRepositoryImpl) FindById(db *gorm.DB, productId string) (domain.Product, error) {
	var product domain.Product

	err := db.Take(&product, "id = ?", productId).Error
	return product, err
}
