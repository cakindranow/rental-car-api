package repository

import (
	"github.com/indrawanagung/food-order-api/model/domain"
	"github.com/indrawanagung/food-order-api/model/web"
	"gorm.io/gorm"
)

type CarRepositoryInterface interface {
	Save(db *gorm.DB, car domain.Car)
	FindAll(db *gorm.DB, name string) []domain.Car
	FindById(db *gorm.DB, carId string) (domain.Car, error)
	CreateOrUpdateOrder(db *gorm.DB, order domain.Order)
	FindAllOrderByUserID(db *gorm.DB, userID string) []web.OrderResponse
	FindAllOrderByAdmin(db *gorm.DB) []web.OrderResponse
	FindOrderByID(db *gorm.DB, orderID string) (domain.Order, error)
	CheckAvailableByOrderDate(db *gorm.DB, startDate string, endDate string, name string) []domain.Car
}
