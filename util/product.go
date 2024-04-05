package util

import (
	"github.com/indrawanagung/food-order-api/model/domain"
	"github.com/indrawanagung/food-order-api/model/web"
)

func ToProductResponse(product domain.Product) web.ProductResponse {
	config := LoadConfig(".")
	return web.ProductResponse{
		Id:       product.ID,
		Name:     product.Name,
		ImageUrl: config.HOST + "/public/images/" + product.ImageUrl,
		Price:    product.Price,
	}
}

func ToListProductResponse(products []domain.Product) []web.ProductResponse {
	//config := LoadConfig(".")

	var productsResponse []web.ProductResponse

	for _, product := range products {
		productsResponse = append(productsResponse, ToProductResponse(product))
	}
	return productsResponse
}
