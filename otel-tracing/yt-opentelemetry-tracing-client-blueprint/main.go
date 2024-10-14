package main

import (
	"context"
	"fmt"
	"github.com/gofiber/contrib/otelfiber/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	"go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.4.0"
	"log"
	"net/http"
)

func initTracer() (*trace.TracerProvider, error) {
	ctx := context.Background()
	exporter, err := otlptracehttp.New(ctx, otlptracehttp.WithEndpoint("localhost:4318"), otlptracehttp.WithInsecure())
	if err != nil {
		return nil, err
	}

	tp := trace.NewTracerProvider(
		trace.WithBatcher(exporter),
		trace.WithSampler(trace.AlwaysSample()),
		trace.WithResource(resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceNameKey.String("service-a"),
		)),
	)
	otel.SetTracerProvider(tp)
	otel.SetTextMapPropagator(propagation.TraceContext{})
	return tp, nil
}

func main() {
	tp, err := initTracer()
	if err != nil {
		log.Fatalf("failed to initialize tracer: %v", err)
	}
	defer func() { _ = tp.Shutdown(context.Background()) }()

	client := &http.Client{
		Transport: otelhttp.NewTransport(http.DefaultTransport),
	}

	app := fiber.New()

	app.Use(otelfiber.Middleware())

	app.Get("/call-service-b", func(c *fiber.Ctx) error {
		tracer := otel.Tracer("service-a")
		ctx, span := tracer.Start(c.UserContext(), "call Order Api")
		defer span.End()

		correlationID := uuid.New().String()

		req, _ := http.NewRequestWithContext(ctx, "GET", "http://localhost:3001/orders/11", nil)
		req.Header.Set("x-correlationid", correlationID)

		// Inject TraceParent to Context
		otel.GetTextMapPropagator().Inject(ctx, propagation.HeaderCarrier(req.Header))

		for k, v := range req.Header {
			fmt.Printf("%s: %s\n", k, v)
		}

		resp, err := client.Do(req)
		if err != nil {
			span.RecordError(err)
			return c.Status(http.StatusInternalServerError).SendString("Failed to call Service B")
		}
		defer resp.Body.Close()

		span.SetAttributes(attribute.String("response_status", resp.Status))

		return c.SendString("Called service B successfully")
	})

	log.Fatal(app.Listen(":3000"))
}
