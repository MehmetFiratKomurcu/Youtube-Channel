package controller

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"strings"
	"yt-swagger-go/src/application/model"
	"yt-swagger-go/src/infra/validation"
)

// GetOrderByCode Getting Order by Code
//
//	@Summary		Getting Order by Code
//	@Description	Getting Order by Code in detail
//	@Tags			Orders
//	@Accept			json
//	@Produce		json
//	@Param			x-correlationid	header		string	true	"code of Order"
//	@Param			orderCode		path		string	true	"code of Order"
//	@Success		200				{string}	string
//	@Router			/orders/code/{orderCode} [get]
func GetOrderByCode(app *fiber.App) fiber.Router {
	return app.Get("/orders/code/:orderCode", func(ctx *fiber.Ctx) error {
		fmt.Printf("Your correlationId is %v", ctx.Locals("correlationId"))

		return ctx.SendString("This is your order Code: " + ctx.Params("orderCode"))
	})
}

// CreateOrder Creating Order
//
//	@Summary		Creating Order
//	@Description	Creating Order with given request
//	@Tags			Orders
//	@Accept			json
//	@Produce		json
//	@Param			x-correlationid	header		string						true	"code of Order"
//	@Param			request			body		model.CreateOrderRequest	true	"Request of Creating Order Object"
//	@Success		200				{string}	string
//	@Failure		400				{string}	string	"Bad Request"
//	@Router			/orders [post]
func CreateOrder(app *fiber.App, customValidator *validation.CustomValidator) fiber.Router {
	return app.Post("/orders", func(ctx *fiber.Ctx) error {
		var request model.CreateOrderRequest
		err := ctx.BodyParser(&request)
		if err != nil {
			return err
		}

		if errs := customValidator.Validate(customValidator.Validator, request); len(errs) > 0 && errs[0].HasError {
			errorMessages := make([]string, 0)

			for _, err2 := range errs {
				errorMessages = append(errorMessages, fmt.Sprintf("%s field has failed. Validation is: %s", err2.Field, err2.Tag))
			}

			return ctx.Status(fiber.StatusBadRequest).JSON(strings.Join(errorMessages, " and that "))
		}

		return ctx.Status(fiber.StatusCreated).JSON("Order created successfully!")
	})
}
