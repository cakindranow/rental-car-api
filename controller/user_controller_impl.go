package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/indrawanagung/food-order-api/exception"
	"github.com/indrawanagung/food-order-api/model/web"
	"github.com/indrawanagung/food-order-api/service"
	"github.com/indrawanagung/food-order-api/util"
)

type UserControllerImpl struct {
	UserService service.UserServiceInterface
}

func NewUserController(userService service.UserServiceInterface) UserControllerInterface {
	return &UserControllerImpl{
		UserService: userService,
	}
}

func (u UserControllerImpl) Save(c *fiber.Ctx) error {
	userCreateRequest := new(web.UserCreateOrUpdateRequest)
	if err := c.BodyParser(userCreateRequest); err != nil {
		panic(exception.NewBadRequestError(err.Error()))
	}

	userID := u.UserService.Save(*userCreateRequest)
	webResponse := web.WebResponse{
		Header: util.HeaderResponseSuccessfully(),
		Data:   map[string]string{"user_id": userID},
	}
	return c.Status(201).JSON(webResponse)
}

func (u UserControllerImpl) Update(c *fiber.Ctx) error {
	userID := c.Params("userID")
	userUpdateRequest := new(web.UserCreateOrUpdateRequest)
	if err := c.BodyParser(userUpdateRequest); err != nil {
		panic(exception.NewBadRequestError(err.Error()))
	}

	u.UserService.Update(userID, *userUpdateRequest)
	webResponse := web.WebResponse{
		Header: util.HeaderResponseSuccessfully(),
		Data:   nil,
	}
	return c.Status(200).JSON(webResponse)
}

func (u UserControllerImpl) GetProfile(c *fiber.Ctx) error {
	userAuth := c.Locals("user").(*jwt.Token)
	claims := userAuth.Claims.(jwt.MapClaims)
	userID := claims["id"].(string)
	userResponse := u.UserService.FindByID(userID)
	webResponse := web.WebResponse{
		Header: util.HeaderResponseSuccessfully(),
		Data:   userResponse,
	}

	return c.Status(200).JSON(webResponse)
}

func (u UserControllerImpl) FindAll(c *fiber.Ctx) error {

	userResponse := u.UserService.FindAll()
	webResponse := web.WebResponse{
		Header: util.HeaderResponseSuccessfully(),
		Data:   userResponse,
	}

	return c.Status(200).JSON(webResponse)
}
