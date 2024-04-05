package service

import (
	"github.com/indrawanagung/food-order-api/model/web"
)

type ProductServiceInterface interface {
	FindAll() []web.ProductResponse
	FindById(id string) web.ProductResponse
	Save(request web.CreateOrUpdateProduct)
}
