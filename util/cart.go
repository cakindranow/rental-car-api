package util

import (
	"github.com/indrawanagung/food-order-api/model/web"
)

func ToCartProductResponse(cart web.ListCartResponse) web.ListCartResponse {
	config := LoadConfig(".")

	return web.ListCartResponse{
		UserID:      cart.UserID,
		ProductID:   cart.ProductID,
		Total:       cart.Total,
		ProductName: cart.ProductName,
		Price:       cart.Price,
		ImageUrl:    config.HOST + "/public/images/" + cart.ImageUrl,
	}
}

func ToCartProductResponses(items []web.ListCartResponse) []web.ListCartResponse {
	var listCartResponses []web.ListCartResponse

	for _, cart := range items {
		listCartResponses = append(listCartResponses, ToCartProductResponse(cart))
	}

	return listCartResponses
}
