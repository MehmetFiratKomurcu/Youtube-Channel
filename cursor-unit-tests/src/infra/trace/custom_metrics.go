package trace

import (
	"context"
	"github.com/gofiber/fiber/v2/log"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/metric"
	"gorm.io/gorm"
	"yt-otel-metrics/src/application/domain/entity"
)

var (
	shippedOrderMetric   metric.Int64ObservableGauge
	cancelledOrderMetric metric.Int64ObservableGauge
)

func InitCustomMetrics(db *gorm.DB) {
	meterProvider := otel.GetMeterProvider()
	meter := meterProvider.Meter("Order Api")

	var err error
	getShippedOrderMetrics(db, err, meter)
	getCancelledOrderMetrics(db, err, meter)
}

func getShippedOrderMetrics(db *gorm.DB, err error, meter metric.Meter) {
	shippedOrderMetric, err = meter.Int64ObservableGauge(
		"shipped_order_metrics",
		metric.WithDescription("Number of shipped_order_metrics"))

	if err != nil {
		log.Fatalf("Failed to initialize shipped_order_metrics: %s", err)
	}

	_, err = meter.RegisterCallback(
		func(ctx context.Context, observer metric.Observer) error {
			var count int64

			result := db.WithContext(ctx).Model(&entity.Order{}).Where("is_shipped = ? AND created_at > now() - interval '7 days'", true).Count(&count)

			if result.Error != nil {
				log.Errorf("Failed to count shipped_order_metrics: %s", result.Error)
			}

			observer.ObserveInt64(shippedOrderMetric, count)
			log.Infof("Recorded metric: %d shipped orders in the last 7 days", count)
			return nil
		}, shippedOrderMetric)

	if err != nil {
		log.Fatalf("Failed to register callback function shipped_order_metrics: %s", err)
	}
}

func getCancelledOrderMetrics(db *gorm.DB, err error, meter metric.Meter) {
	cancelledOrderMetric, err = meter.Int64ObservableGauge(
		"cancelled_order_metrics",
		metric.WithDescription("Number of cancelled_order_metrics"))

	if err != nil {
		log.Fatalf("Failed to initialize cancelled_order_metrics: %s", err)
	}

	_, err = meter.RegisterCallback(
		func(ctx context.Context, observer metric.Observer) error {
			var count int64

			result := db.WithContext(ctx).Model(&entity.Order{}).Where("is_cancelled = ? AND created_at > now() - interval '7 days'", true).Count(&count)

			if result.Error != nil {
				log.Errorf("Failed to count shipped_order_metrics: %s", result.Error)
			}

			observer.ObserveInt64(cancelledOrderMetric, count)
			log.Infof("Recorded metric: %d cancelled orders in the last 7 days", count)
			return nil
		}, cancelledOrderMetric)

	if err != nil {
		log.Fatalf("Failed to register callback function cancel_order_metrics: %s", err)
	}
}
