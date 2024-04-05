package controller

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/indrawanagung/food-order-api/exception"
	"github.com/indrawanagung/food-order-api/model/web"
	"github.com/indrawanagung/food-order-api/service"
	"github.com/indrawanagung/food-order-api/util"
)

type ProductControllerImpl struct {
	ProductService service.ProductServiceInterface
}

func NewProductController(productService service.ProductServiceInterface) ProductControllerInterface {
	return &ProductControllerImpl{
		ProductService: productService,
	}
}

func (c ProductControllerImpl) FindAll(ctx *fiber.Ctx) error {
	products := c.ProductService.FindAll()
	webResponse := web.WebResponse{
		Header: util.HeaderResponseSuccessfully(),
		Data:   products,
	}
	ctx.Accepts("Access-Control-Allow-Headers", "*1")
	return ctx.Status(200).JSON(webResponse)

}

func (c ProductControllerImpl) FindByID(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	products := c.ProductService.FindById(id)
	webResponse := web.WebResponse{
		Header: util.HeaderResponseSuccessfully(),
		Data:   products,
	}
	return ctx.Status(200).JSON(webResponse)
}

func (c ProductControllerImpl) Save(ctx *fiber.Ctx) error {
	request := new(web.CreateOrUpdateProduct)
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
	c.ProductService.Save(*request)

	webResponse := web.WebResponse{
		Header: util.HeaderResponseSuccessfully(),
		Data:   nil,
	}
	return ctx.Status(201).JSON(webResponse)
}
