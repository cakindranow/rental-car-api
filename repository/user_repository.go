package repository

import (
	"github.com/indrawanagung/food-order-api/model/domain"
	"gorm.io/gorm"
)

type UserRepositoryInterface interface {
	SaveOrUpdate(tx *gorm.DB, user domain.User) error
	FindByID(tx *gorm.DB, userID string) (error, domain.User)
	FindByEmail(tx *gorm.DB, email string) (error, domain.User)
	FindAll(tx *gorm.DB) []domain.User
}
