package controller

import "github.com/gofiber/fiber/v2"

type AuthControllerInterface interface {
	Login(ctx *fiber.Ctx) error
}
