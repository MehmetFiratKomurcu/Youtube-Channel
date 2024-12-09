package main

import (
	"context"
	"github.com/ansrivas/fiberprometheus/v2"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/swagger"
	"github.com/uptrace/opentelemetry-go-extra/otelgorm"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	_ "yt-otel-metrics/docs"
	"yt-otel-metrics/src/application/controller"
	"yt-otel-metrics/src/application/domain/persistance"
	"yt-otel-metrics/src/application/domain/services"
	"yt-otel-metrics/src/infra/middleware"
	"yt-otel-metrics/src/infra/trace"
	"yt-otel-metrics/src/infra/validation"
)

var validate = validator.New()

// @title			Order Api
// @version		1.0
// @description	This is an Order Api just for young people
// @termsOfService	http://swagger.io/terms/
func main() {
	tp, err := trace.InitTracer()
	if err != nil {
		log.Fatal("something is wrong")
	}
	defer func() {
		_ = tp.Shutdown(context.Background())
	}()

	mp, err := trace.InitMeter()
	if err != nil {
		log.Fatal("something is wrong")
	}
	defer func() {
		_ = mp.Shutdown(context.Background())
	}()

	app := fiber.New()

	prometheus := fiberprometheus.New("order-api")
	prometheus.RegisterAt(app, "/metrics")

	app.Use(prometheus.Middleware)

	app.Get("/swagger/*", swagger.HandlerDefault)

	connString := "host=localhost user=postgres password=mysecretpassword dbname=postgres port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(connString), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	err = db.Use(otelgorm.NewPlugin())

	if err != nil {
		panic("aaaağğhh")
	}

	trace.InitMetrics(db)

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
	controller.ShipOrderByCargoCode(app, orderService)

	app.Listen(":3001")
}
