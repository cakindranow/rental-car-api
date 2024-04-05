package exception

import (
	"errors"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/indrawanagung/food-order-api/model/web"
)

var ErrNotFound = errors.New("Resource was not found")

func ErrorHandler(ctx *fiber.Ctx, err error) error {
	_, ok := err.(validator.ValidationErrors)
	if ok {
		return ctx.Status(400).JSON(web.WebResponse{
			Header: web.Header{
				Message: err.Error(),
				Error:   true,
			},
			Data: nil,
		})
	}

	_, ok = err.(NotFoundError)
	if ok {
		return ctx.Status(404).JSON(web.WebResponse{
			Header: web.Header{
				Message: err.Error(),
				Error:   true,
			},
			Data: nil,
		})
	}

	_, ok = err.(BadRequestError)
	if ok {
		return ctx.Status(400).JSON(web.WebResponse{
			Header: web.Header{
				Message: err.Error(),
				Error:   true,
			},
			Data: nil,
		})
	}

	_, ok = err.(UnauthorizedError)
	if ok {
		return ctx.Status(401).JSON(web.WebResponse{
			Header: web.Header{
				Message: err.Error(),
				Error:   true,
			},
			Data: nil,
		})
	}

	_, ok = err.(ForbiddenError)
	if ok {
		return ctx.Status(403).JSON(web.WebResponse{
			Header: web.Header{
				Message: err.Error(),
				Error:   true,
			},
			Data: nil,
		})
	}

	code := fiber.StatusInternalServerError

	// Retrieve the custom status code if it's a *fiber.Error
	var e *fiber.Error
	if errors.As(err, &e) {
		code = e.Code
	}

	// Set Content-Type: text/plain; charset=utf-8
	ctx.Set(fiber.HeaderContentType, fiber.MIMETextPlainCharsetUTF8)

	// Return status code with error message
	return ctx.Status(code).JSON(web.WebResponse{
		Header: web.Header{
			Message: err.Error(),
			Error:   true,
		},
		Data: nil,
	})
}
