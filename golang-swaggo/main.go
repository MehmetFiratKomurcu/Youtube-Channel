package main

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/swagger"
	_ "yt-swagger-go/docs"
	"yt-swagger-go/src/application/controller"
	"yt-swagger-go/src/infra/middleware"
	"yt-swagger-go/src/infra/validation"
)

var validate = validator.New()

//	@title			Order Api
//	@version		1.0
//	@description	This is an Order Api just for young people
//	@termsOfService	http://swagger.io/terms/
func main() {
	app := fiber.New()
	app.Get("/swagger/*", swagger.HandlerDefault)

	customValidator := &validation.CustomValidator{
		Validator: validate,
	}

	err := validation.AddOldAgeCustomValidation(customValidator)
	if err != nil {
		return
	}

	app.Use(recover.New())
	app.Use(cors.New())

	// middleware
	middleware.AddCorrelationId(app)

	// endpoints
	controller.GetOrderByCode(app)
	controller.CreateOrder(app, customValidator)

	app.Listen(":3000")
}
