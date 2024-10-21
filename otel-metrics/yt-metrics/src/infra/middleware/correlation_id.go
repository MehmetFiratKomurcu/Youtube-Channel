package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func AddCorrelationId(app *fiber.App) fiber.Router {
	return app.Use("/orders", func(ctx *fiber.Ctx) error {
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
}
