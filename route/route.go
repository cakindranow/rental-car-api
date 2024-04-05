package route

import (
	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/indrawanagung/food-order-api/controller"
	"github.com/indrawanagung/food-order-api/exception"
	"github.com/spf13/viper"
)

func New(
	ProductController controller.ProductControllerInterface,
	UserController controller.UserControllerInterface,
	AuthController controller.AuthControllerInterface,
	CartController controller.CartControllerInterface,
	CarController controller.CarControllerInterface,
) *fiber.App {
	authentication := jwtware.New(jwtware.Config{
		SigningKey: jwtware.SigningKey{Key: []byte(viper.GetString("SECRET_KEY"))},
	})
	app := fiber.New(fiber.Config{ErrorHandler: exception.ErrorHandler})

	app.Use(cors.New())
	app.Use(recover.New())

	api := app.Group("/api")

	api.Post("/auth/login", AuthController.Login)

	api.Get("/users/all", authentication, UserController.FindAll)
	api.Get("/users", authentication, UserController.GetProfile)
	api.Post("/users", UserController.Save)
	api.Put("/users/:userID", authentication, UserController.Update)

	api.Get("/products", authentication, ProductController.FindAll)
	api.Get("/products/:id", authentication, ProductController.FindByID)
	api.Post("/products", authentication, ProductController.Save)

	//api.Get("/cars", authentication, ProductController.FindAll)
	//api.Get("/cars/:id", authentication, ProductController.FindByID)
	//api.Post("/cars", authentication, ProductController.Save)
	//

	api.Get("/cars", authentication, CarController.FindAll)
	api.Get("/cars/:id", authentication, CarController.FindByID)
	api.Post("/cars", authentication, CarController.Save)

	api.Get("/orders", authentication, CarController.FindAllOrder)
	api.Post("/orders", authentication, CarController.CreateOrder)
	api.Get("/orders/cancel/:id", authentication, CarController.CanceledOrderByUserID)
	api.Get("/orders/approve/:id", authentication, CarController.ApproveCarOrderByAdmin)
	api.Get("/orders/reject/:id", authentication, CarController.RejectCarOrderByAdmin)
	api.Get("/orders/return/:id", authentication, CarController.ReturnedCarOrderByUserID)

	return app
}
