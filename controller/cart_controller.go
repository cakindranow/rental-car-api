package controller

import "github.com/gofiber/fiber/v2"

type CartControllerInterface interface {
	FindAll(ctx *fiber.Ctx) error
	FindByProductID(ctx *fiber.Ctx) error
	Save(ctx *fiber.Ctx) error
	Delete(ctx *fiber.Ctx) error
}
