package util

import (
	"github.com/indrawanagung/food-order-api/model/domain"
	"github.com/indrawanagung/food-order-api/model/web"
)

func ToCarResponse(car domain.Car) web.CarResponse {
	config := LoadConfig(".")
	return web.CarResponse{
		Id:              car.ID,
		Brand:           car.Brand,
		Model:           car.Model,
		Plat:            car.Plat,
		DailyRentalRate: car.DailyRentalRate,
		Desc:            car.Desc,
		ImageUrl:        config.HOST + "/public/images/" + car.ImageUrl,
	}
}

func ToListCarResponse(cars []domain.Car) []web.CarResponse {
	//config := LoadConfig(".")

	var carsResponse []web.CarResponse

	for _, car := range cars {
		carsResponse = append(carsResponse, ToCarResponse(car))
	}
	return carsResponse
}
