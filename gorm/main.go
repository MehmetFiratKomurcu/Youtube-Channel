package main

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/swagger"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	_ "yt-gorm/docs"
	"yt-gorm/src/application/controller"
	"yt-gorm/src/application/domain/persistance"
	"yt-gorm/src/application/domain/services"
	"yt-gorm/src/infra/middleware"
	"yt-gorm/src/infra/validation"
)

var validate = validator.New()

// @title			Order Api
// @version		1.0
// @description	This is an Order Api just for young people
// @termsOfService	http://swagger.io/terms/
func main() {
	app := fiber.New()
	app.Get("/swagger/*", swagger.HandlerDefault)

	connString := "host=localhost user=postgres password=mysecretpassword dbname=postgres port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(connString), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	customValidator := &validation.CustomValidator{
		Validator: validate,
	}

	app.Use(recover.New())
	app.Use(cors.New())

	// middleware
	middleware.AddCorrelationId(app)

	// repositories
	orderRepository := persistance.NewOrderRepository(db)

	// services
	orderService := services.NewOrderService(orderRepository)

	// endpoints
	controller.GetOrderById(app, orderService)
	controller.CreateOrder(app, customValidator, orderService)

	app.Listen(":3001")
}
