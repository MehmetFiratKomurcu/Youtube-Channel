package controller

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/propagation"
	"strconv"
	"strings"
	"yt-otel-metrics/src/application/domain/services"
	"yt-otel-metrics/src/application/model"
	"yt-otel-metrics/src/infra/validation"
)

// GetOrderById Getting Order by Code
//
//	@Summary		Getting Order by Id
//	@Description	Getting Order by Id in detail
//	@Tags			Orders
//	@Accept			json
//	@Produce		json
//	@Param			x-correlationid	header		string	true	"id of Order"
//	@Param			id				path		string	true	"id of Order"
//	@Failure		400				{string}	string
//	@Failure		404				{string}	string
//	@Success		200				{string}	string
//	@Router			/orders/{id} [get]
func GetOrderById(app *fiber.App, orderService services.OrderService) fiber.Router {
	return app.Get("/orders/:id", func(ctx *fiber.Ctx) error {
		propagator := otel.GetTextMapPropagator()
		extractCtx := propagator.Extract(ctx.UserContext(), propagation.HeaderCarrier(ctx.GetReqHeaders()))

		fmt.Printf("Your correlationId is %v", ctx.Locals("correlationId"))

		tracer := otel.Tracer("Order Api")
		userContext, span := tracer.Start(extractCtx, "GetOrderById Controller")
		defer span.End()

		traceParent := ctx.Get("traceparent")
		fmt.Println("traceParent: " + traceParent)

		orderId := ctx.Params("id")
		id, err := strconv.ParseInt(orderId, 10, 64)

		if err != nil {
			return ctx.Status(fiber.StatusBadRequest).JSON("Order id is not valid")
		}

		order := orderService.GetOrderById(userContext, id)

		span.SetAttributes(attribute.String("order.id", orderId))

		if order == nil {
			span.SetAttributes(attribute.String("error", "Order not found"))
			span.SetStatus(codes.Error, "Order not found")
			return ctx.Status(fiber.StatusNotFound).JSON("Order Not Found, sorry! :(")
		}

		return ctx.Status(fiber.StatusOK).JSON(order)
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
//	@Param			request			body		model.CreateOrderCommand	true	"Request of Creating Order Object"
//	@Success		201				{string}	string
//	@Failure		400				{string}	string	"Bad Request"
//	@Router			/orders [post]
func CreateOrder(app *fiber.App, customValidator *validation.CustomValidator, orderService services.OrderService) fiber.Router {
	return app.Post("/orders", func(ctx *fiber.Ctx) error {
		var createOrderCommand model.CreateOrderCommand
		err := ctx.BodyParser(&createOrderCommand)
		if err != nil {
			return err
		}

		err2, hasError := validateCreateOrderRequest(ctx, customValidator, createOrderCommand)
		if hasError {
			return err2
		}

		createdOrder := orderService.CreateOrder(createOrderCommand)

		return ctx.Status(fiber.StatusCreated).JSON(createdOrder)
	})
}

// ShipOrderByCargoCode Ship Order by CargoCode
//
//	@Summary		Shipping Order
//	@Description	Ship Order with given request
//	@Tags			Orders
//	@Accept			json
//	@Produce		json
//	@Param			x-correlationid	header		string	true	"code of Order"
//	@Param			cargoCode		path		string	true	"Request for cargo code or orders"
//	@Success		200				{string}	string
//	@Failure		400				{string}	string	"Bad Request"
//	@Router			/orders/cargo-code/{cargoCode}/ship [post]
func ShipOrderByCargoCode(app *fiber.App, orderService services.OrderService) fiber.Router {
	return app.Post("/orders/cargo-code/:cargoCode/ship", func(ctx *fiber.Ctx) error {
		cargoCode := ctx.Params("cargoCode")

		if cargoCode == "" {
			return ctx.Status(fiber.StatusBadRequest).JSON("invalid cargo code")
		}

		err := orderService.ShipOrderByCargoCode(cargoCode)

		if err != nil {
			return ctx.Status(fiber.StatusBadRequest).JSON(err.Error())
		}

		return ctx.Status(fiber.StatusOK).JSON("Orders ship successfully! yayyyy!")
	})
}

func validateCreateOrderRequest(ctx *fiber.Ctx, customValidator *validation.CustomValidator, request model.CreateOrderCommand) (error, bool) {
	if errs := customValidator.Validate(customValidator.Validator, request); len(errs) > 0 && errs[0].HasError {
		errorMessages := make([]string, 0)

		for _, err2 := range errs {
			errorMessages = append(errorMessages, fmt.Sprintf("%s field has failed. Validation is: %s", err2.Field, err2.Tag))
		}

		return ctx.Status(fiber.StatusBadRequest).JSON(strings.Join(errorMessages, " and that ")), true
	}
	return nil, false
}
