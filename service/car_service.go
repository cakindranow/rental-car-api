package service

import (
	"github.com/indrawanagung/food-order-api/model/web"
)

type CarServiceInterface interface {
	FindAll(startDate string, endDate string, name string) []web.CarResponse
	FindById(id string) web.CarResponse
	Save(request web.CreateOrUpdateCar)
	CreateOrder(request web.CreateOrderRequest, userID string)
	FindAllOrder(userID string) []web.OrderResponse
	CanceledOrderByUserID(userID string, orderID string)
	ReturnedCarOrderByUserID(userID string, orderID string, plat string)
	ApproveCarOrderByAdmin(userID string, orderID string)
	RejectCarOrderByAdmin(userID string, orderID string, noteAdmin string)
}
