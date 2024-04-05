package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/indrawanagung/food-order-api/exception"
	"github.com/indrawanagung/food-order-api/model/web"
	"github.com/indrawanagung/food-order-api/service"
	"github.com/indrawanagung/food-order-api/util"
)

type AuthControllerImpl struct {
	AuthService service.AuthServiceInterface
}

func NewAuthController(authService service.AuthServiceInterface) AuthControllerInterface {
	return &AuthControllerImpl{AuthService: authService}
}

func (c AuthControllerImpl) Login(ctx *fiber.Ctx) error {
	loginRequest := new(web.LoginRequest)
	if err := ctx.BodyParser(loginRequest); err != nil {
		panic(exception.NewBadRequestError(err.Error()))
	}
	loginResponse := c.AuthService.Login(*loginRequest)
	webResponse := web.WebResponse{
		Header: util.HeaderResponseSuccessfully(),
		Data:   loginResponse,
	}
	return ctx.Status(200).JSON(webResponse)

}
