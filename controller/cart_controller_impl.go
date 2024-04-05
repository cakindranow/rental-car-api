package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/indrawanagung/food-order-api/exception"
	"github.com/indrawanagung/food-order-api/model/web"
	"github.com/indrawanagung/food-order-api/service"
	"github.com/indrawanagung/food-order-api/util"
)

type CartControllerImpl struct {
	CartService service.CartServiceInterface
}

func NewCartController(cartService service.CartServiceInterface) CartControllerInterface {
	return &CartControllerImpl{CartService: cartService}
}

func (c CartControllerImpl) FindAll(ctx *fiber.Ctx) error {
	userAuth := ctx.Locals("user").(*jwt.Token)
	claims := userAuth.Claims.(jwt.MapClaims)
	userID := claims["id"].(string)

	cartResponses := c.CartService.FindAll(userID)
	webResponse := web.WebResponse{
		Header: util.HeaderResponseSuccessfully(),
		Data:   cartResponses,
	}
	return ctx.Status(200).JSON(webResponse)
}

func (c CartControllerImpl) FindByProductID(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	userAuth := ctx.Locals("user").(*jwt.Token)
	claims := userAuth.Claims.(jwt.MapClaims)
	userID := claims["id"].(string)

	cartResponse := c.CartService.FindByProductAndUserID(id, userID)

	webResponse := web.WebResponse{
		Header: util.HeaderResponseSuccessfully(),
		Data:   cartResponse,
	}
	return ctx.Status(200).JSON(webResponse)
}

func (c CartControllerImpl) Save(ctx *fiber.Ctx) error {
	userAuth := ctx.Locals("user").(*jwt.Token)
	claims := userAuth.Claims.(jwt.MapClaims)
	userID := claims["id"].(string)

	request := new(web.CartCreateRequest)
	request.UserID = userID

	if err := ctx.BodyParser(request); err != nil {
		panic(exception.NewBadRequestError(err.Error()))
	}

	c.CartService.Save(*request)
	webResponse := web.WebResponse{
		Header: util.HeaderResponseSuccessfully(),
		Data:   nil,
	}
	return ctx.Status(200).JSON(webResponse)
}

func (c CartControllerImpl) Delete(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	userAuth := ctx.Locals("user").(*jwt.Token)
	claims := userAuth.Claims.(jwt.MapClaims)
	userID := claims["id"].(string)

	c.CartService.Delete(id, userID)

	webResponse := web.WebResponse{
		Header: util.HeaderResponseSuccessfully(),
		Data:   nil,
	}
	return ctx.Status(200).JSON(webResponse)
}
