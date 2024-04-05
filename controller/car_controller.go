package controller

import (
	"github.com/gofiber/fiber/v2"
)

type CarControllerInterface interface {
	FindAll(ctx *fiber.Ctx) error
	FindByID(ctx *fiber.Ctx) error
	Save(ctx *fiber.Ctx) error
	CreateOrder(ctx *fiber.Ctx) error
	FindAllOrder(ctx *fiber.Ctx) error
	CanceledOrderByUserID(ctx *fiber.Ctx) error
	ReturnedCarOrderByUserID(ctx *fiber.Ctx) error
	ApproveCarOrderByAdmin(ctx *fiber.Ctx) error
	RejectCarOrderByAdmin(ctx *fiber.Ctx) error
}
