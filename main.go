package main

import (
	"crypto/tls"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2/log"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/indrawanagung/food-order-api/controller"
	"github.com/indrawanagung/food-order-api/db"
	"github.com/indrawanagung/food-order-api/repository"
	"github.com/indrawanagung/food-order-api/route"
	"github.com/indrawanagung/food-order-api/service"
	"github.com/indrawanagung/food-order-api/util"
)

func main() {
	validate := validator.New()

	config := util.LoadConfig(".")
	database := db.OpenConnection(config.DBSource)

	productRepository := repository.NewProductRepository()
	userRepository := repository.NewUserRepository()
	cartRepository := repository.NewCartRepository()
	carRepository := repository.NewCarRepository()

	carService := service.NewCarService(database, validate, carRepository, userRepository)
	productService := service.NewProductService(database, validate, productRepository)
	userService := service.NewUserService(database, userRepository, validate)
	authService := service.NewAuthService(userRepository, database, validate)
	cartService := service.NewCartService(cartRepository, database, validate)

	carController := controller.NewCarController(carService)
	authController := controller.NewAuthController(authService)
	userController := controller.NewUserController(userService)
	productController := controller.NewProductController(productService)
	cartController := controller.NewCartController(cartService)

	app := route.New(productController, userController, authController, cartController, carController)
	app.Static("/api/public/images", "./public/images")
	app.Use(logger.New(logger.Config{
		Format:     "${cyan}[${time}] ${white}${pid} ${red}${status} ${blue}[${method}] ${white}${path}\n",
		TimeFormat: "02-Jan-2006",
		TimeZone:   "UTC",
	}))

	// Create tls certificate
	cer, err := tls.LoadX509KeyPair("certs/ssl.cert", "certs/ssl.key")
	if err != nil {
		log.Fatal(err)
	}
	config2 := &tls.Config{Certificates: []tls.Certificate{cer}}

	// Create custom listener
	ln, err := tls.Listen("tcp", ":4000", config2)
	if err != nil {
		panic(err)
	}

	// Start server with https/ssl enabled on http://localhost:443
	log.Fatal(app.Listener(ln))

	//f, _ := os.OpenFile("test.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	//log.SetOutput(f)
	//log.Info("server running on port 4000")
	//log.Fatal(app.Listen(":4000"))
}
