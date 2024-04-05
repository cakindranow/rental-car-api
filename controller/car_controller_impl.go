package controller

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/indrawanagung/food-order-api/exception"
	"github.com/indrawanagung/food-order-api/model/web"
	"github.com/indrawanagung/food-order-api/service"
	"github.com/indrawanagung/food-order-api/util"
)

type CarControllerImpl struct {
	CarService service.CarServiceInterface
}

func NewCarController(carService service.CarServiceInterface) CarControllerInterface {
	return &CarControllerImpl{
		CarService: carService,
	}
}

func (c CarControllerImpl) FindAll(ctx *fiber.Ctx) error {
	startDate := ctx.Query("start_date")
	endDate := ctx.Query("end_date")
	name := ctx.Query("name")

	carts := c.CarService.FindAll(startDate, endDate, name)
	webResponse := web.WebResponse{
		Header: util.HeaderResponseSuccessfully(),
		Data:   carts,
	}
	ctx.Accepts("Access-Control-Allow-Headers", "*1")
	return ctx.Status(200).JSON(webResponse)

}

func (c CarControllerImpl) FindByID(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	carts := c.CarService.FindById(id)
	webResponse := web.WebResponse{
		Header: util.HeaderResponseSuccessfully(),
		Data:   carts,
	}
	return ctx.Status(200).JSON(webResponse)
}

func (c CarControllerImpl) Save(ctx *fiber.Ctx) error {
	request := new(web.CreateOrUpdateCar)
	if err := ctx.BodyParser(request); err != nil {
		panic(exception.NewBadRequestError(err.Error()))
	}

	file, err := ctx.FormFile("image")
	if err != nil {
		panic(exception.NewBadRequestError("file image is not found"))
	}
	fileName := uuid.NewString() + ".jpg"
	// Save file to root directory:
	err = ctx.SaveFile(file, fmt.Sprintf("./public/images/%s", fileName))
	if err != nil {
		panic(err)
	}

	request.ImageUrl = fileName
	c.CarService.Save(*request)

	webResponse := web.WebResponse{
		Header: util.HeaderResponseSuccessfully(),
		Data:   nil,
	}
	return ctx.Status(201).JSON(webResponse)
}

func (c CarControllerImpl) CreateOrder(ctx *fiber.Ctx) error {
	request := new(web.CreateOrderRequest)
	if err := ctx.BodyParser(request); err != nil {
		panic(exception.NewBadRequestError(err.Error()))
	}

	userAuth := ctx.Locals("user").(*jwt.Token)
	claims := userAuth.Claims.(jwt.MapClaims)
	userID := claims["id"].(string)
	fmt.Println(userID)
	c.CarService.CreateOrder(*request, userID)

	webResponse := web.WebResponse{
		Header: util.HeaderResponseSuccessfully(),
		Data:   nil,
	}
	return ctx.Status(201).JSON(webResponse)
}

func (c CarControllerImpl) FindAllOrder(ctx *fiber.Ctx) error {

	userAuth := ctx.Locals("user").(*jwt.Token)
	claims := userAuth.Claims.(jwt.MapClaims)
	userID := claims["id"].(string)

	ordersResonse := c.CarService.FindAllOrder(userID)

	webResponse := web.WebResponse{
		Header: util.HeaderResponseSuccessfully(),
		Data:   ordersResonse,
	}
	return ctx.Status(200).JSON(webResponse)
}

func (c CarControllerImpl) CanceledOrderByUserID(ctx *fiber.Ctx) error {
	userAuth := ctx.Locals("user").(*jwt.Token)
	claims := userAuth.Claims.(jwt.MapClaims)
	userID := claims["id"].(string)
	orderID := ctx.Params("id")

	c.CarService.CanceledOrderByUserID(userID, orderID)

	webResponse := web.WebResponse{
		Header: util.HeaderResponseSuccessfully(),
		Data:   nil,
	}
	return ctx.Status(200).JSON(webResponse)
}

func (c CarControllerImpl) ReturnedCarOrderByUserID(ctx *fiber.Ctx) error {
	userAuth := ctx.Locals("user").(*jwt.Token)
	claims := userAuth.Claims.(jwt.MapClaims)
	userID := claims["id"].(string)
	orderID := ctx.Params("id")
	plat := ctx.Query("plat")

	c.CarService.ReturnedCarOrderByUserID(userID, orderID, plat)

	webResponse := web.WebResponse{
		Header: util.HeaderResponseSuccessfully(),
		Data:   nil,
	}
	return ctx.Status(200).JSON(webResponse)
}

func (c CarControllerImpl) ApproveCarOrderByAdmin(ctx *fiber.Ctx) error {
	userAuth := ctx.Locals("user").(*jwt.Token)
	claims := userAuth.Claims.(jwt.MapClaims)
	userID := claims["id"].(string)
	orderID := ctx.Params("id")

	c.CarService.ApproveCarOrderByAdmin(userID, orderID)

	webResponse := web.WebResponse{
		Header: util.HeaderResponseSuccessfully(),
		Data:   nil,
	}
	return ctx.Status(200).JSON(webResponse)
}

func (c CarControllerImpl) RejectCarOrderByAdmin(ctx *fiber.Ctx) error {
	userAuth := ctx.Locals("user").(*jwt.Token)
	claims := userAuth.Claims.(jwt.MapClaims)
	userID := claims["id"].(string)
	orderID := ctx.Params("id")
	note := ctx.Query("note")

	c.CarService.RejectCarOrderByAdmin(userID, orderID, note)

	webResponse := web.WebResponse{
		Header: util.HeaderResponseSuccessfully(),
		Data:   nil,
	}
	return ctx.Status(200).JSON(webResponse)
}
