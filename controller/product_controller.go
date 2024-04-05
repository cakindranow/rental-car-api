package controller

import "github.com/gofiber/fiber/v2"

type ProductControllerInterface interface {
	FindAll(ctx *fiber.Ctx) error
	FindByID(ctx *fiber.Ctx) error
	Save(ctx *fiber.Ctx) error
}
