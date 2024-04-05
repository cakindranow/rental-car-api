package repository

import (
	"errors"
	"fmt"
	"github.com/gofiber/fiber/v2/log"
	"github.com/indrawanagung/food-order-api/model/domain"
	"github.com/indrawanagung/food-order-api/model/web"
	"gorm.io/gorm"
)

type CartRepositoryImpl struct {
}

func NewCartRepository() CartRepositoryInterface {
	return &CartRepositoryImpl{}
}

func (c CartRepositoryImpl) FindAll(db *gorm.DB, userID string) []web.ListCartResponse {
	var carts []web.ListCartResponse
	err := db.Raw("SELECT c.user_id, c.product_id, c.total , p.\"name\" "+
		"as product_name, p.price, p.image_url \n"+
		"FROM carts c inner join products p on p.id = c.product_id \n where c.user_id = ? order by p.id desc", userID).Scan(&carts).Error
	if err != nil {
		log.Error(err)
		panic(err)
	}
	return carts
}

func (c CartRepositoryImpl) FindByProductAndUserID(db *gorm.DB, productID string, userID string) (domain.Cart, error) {
	var cart domain.Cart

	err := db.Take(&cart, "product_id = ? and user_id = ? ", productID, userID).Error

	if err != nil {
		if err.Error() != "record not found" {
			log.Error(err)
			panic(err)
		}
		return cart, errors.New(fmt.Sprintf("cart product is not found"))
	}

	return cart, nil
}

func (c CartRepositoryImpl) Save(db *gorm.DB, request domain.Cart) {
	err := db.Save(&request).Error
	if err != nil {
		log.Error(err)
		panic(err)
	}
}

func (c CartRepositoryImpl) Delete(db *gorm.DB, productID string, userID string) {
	cartProduct := domain.Cart{}
	err := db.Delete(&cartProduct, "product_id = ? and user_id = ?", productID, userID).Error
	if err != nil {
		log.Error(err)
		panic(err)
	}
}
