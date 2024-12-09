package trace

import (
	"context"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/metric"
	"gorm.io/gorm"
	"log"
	"yt-otel-metrics/src/application/domain/entity"
)

var (
	shippedOrdersMetric   metric.Int64ObservableGauge
	cancelledOrdersMetric metric.Int64ObservableGauge
)

func InitMetrics(db *gorm.DB) {
	meterProvider := otel.GetMeterProvider()
	meter := meterProvider.Meter("Order Api")

	var err error
	err = initShippedOrdersMetric(db, err, meter)
	collectCancelledOrdersMetric(db, err, meter)
}

func initShippedOrdersMetric(db *gorm.DB, err error, meter metric.Meter) error {
	shippedOrdersMetric, err = meter.Int64ObservableGauge(
		"shipped_orders_count_metric",
		metric.WithDescription("Counts the number of shipped orders"),
	)

	if err != nil {
		log.Fatalf("Failed to create metric: %v", err)
	}

	_, err = meter.RegisterCallback(
		func(ctx context.Context, observer metric.Observer) error {
			var count int64

			result := db.WithContext(ctx).Model(&entity.Order{}).
				Where("is_shipped = ? AND created_at > NOW() - INTERVAL '7 DAYS'", true).
				Count(&count)

			if result.Error != nil {
				log.Printf("Error querying database: %v", result.Error)
			}

			observer.ObserveInt64(shippedOrdersMetric, count)
			log.Printf("Recorded metric: %d shipped orders in the last 7 days", count)

			return nil
		}, shippedOrdersMetric)

	if err != nil {
		log.Fatalf("Failed to register callback: %v", err)
	}
	return err
}

func collectCancelledOrdersMetric(db *gorm.DB, err error, meter metric.Meter) {
	cancelledOrdersMetric, err = meter.Int64ObservableGauge(
		"cancelled_orders_count_metric",
		metric.WithDescription("Counts the number of cancelled orders"),
	)
	if err != nil {
		log.Fatalf("Failed to create metric: %v", err)
	}

	_, err = meter.RegisterCallback(
		func(ctx context.Context, observer metric.Observer) error {
			var count int64

			result := db.WithContext(ctx).Model(&entity.Order{}).
				Where("is_cancelled = ? AND created_at > NOW() - INTERVAL '7 DAYS'", true).
				Count(&count)

			if result.Error != nil {
				log.Printf("Error querying database: %v", result.Error)
			}

			observer.ObserveInt64(cancelledOrdersMetric, count)
			log.Printf("Recorded metric: %d shipped orders in the last 7 days", count)

			return nil
		}, cancelledOrdersMetric)

	if err != nil {
		log.Fatalf("Failed to register callback: %v", err)
	}
}
