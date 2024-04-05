package service

import "github.com/indrawanagung/food-order-api/model/web"

type AuthServiceInterface interface {
	Login(request web.LoginRequest) web.LoginResponse
}
