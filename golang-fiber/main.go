package main

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/google/uuid"
	"strings"
)

type CreateOrderRequest struct {
	ShipmentNumber string `json:"shipmentNumber" validate:"required"`
	CountryCode    string `json:"countryCode" validate:"required,len=2"`
	Age            int    `json:"age" validate:"required,oldAge"`
}

type ValidationError struct {
	HasError bool
	Field    string
	Tag      string
	Value    interface{}
}

type CustomValidator struct {
	validator *validator.Validate
}

var validate = validator.New()

func (v CustomValidator) Validate(data interface{}) []ValidationError {
	var validationErrors []ValidationError

	errs := validate.Struct(data)
	if errs != nil {
		for _, err := range errs.(validator.ValidationErrors) {
			var ve ValidationError

			ve.Value = err.Value()
			ve.Field = err.Field()
			ve.Tag = err.Tag()
			ve.HasError = true

			validationErrors = append(validationErrors, ve)
		}
	}

	return validationErrors
}

func main() {
	app := fiber.New()

	customValidator := &CustomValidator{
		validator: validate,
	}

	customValidator.validator.RegisterValidation("oldAge", func(fl validator.FieldLevel) bool {
		return fl.Field().Int() < 40
	})

	app.Use(recover.New())

	app.Use(func(ctx *fiber.Ctx) error {
		//panic("OMG I'have something!")
		fmt.Println("You have called " + string(ctx.Request().RequestURI()))

		return ctx.Next()
	})

	app.Use("/orders/code/:orderCode", func(ctx *fiber.Ctx) error {
		correlationId := ctx.Get("x-correlationid")

		if correlationId == "" {
			return ctx.Status(fiber.StatusBadRequest).JSON("CorrelationId is mandatory")
		}

		_, err := uuid.Parse(correlationId)

		if err != nil {
			return ctx.Status(fiber.StatusBadRequest).JSON("CorrelationId is not a guid")
		}

		ctx.Locals("correlationId", correlationId)
		return ctx.Next()
	})

	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.SendString("Hey! I'm here!")
	})

	app.Get("/orders/code/:orderCode", func(ctx *fiber.Ctx) error {
		fmt.Printf("Your correlationId is %v", ctx.Locals("correlationId"))

		return ctx.SendString("This is your order Code: " + ctx.Params("orderCode"))
	})

	app.Post("/orders", func(ctx *fiber.Ctx) error {
		var request CreateOrderRequest
		err := ctx.BodyParser(&request)
		if err != nil {
			return err
		}

		if errs := customValidator.Validate(request); len(errs) > 0 && errs[0].HasError {
			errorMessages := make([]string, 0)

			for _, err2 := range errs {
				errorMessages = append(errorMessages, fmt.Sprintf("%s field has failed. Validation is: %s", err2.Field, err2.Tag))
			}

			return ctx.Status(fiber.StatusBadRequest).JSON(strings.Join(errorMessages, " and that "))
		}

		return ctx.Status(fiber.StatusCreated).JSON("Order created successfully!")
	})

	app.Listen(":3000")
}
