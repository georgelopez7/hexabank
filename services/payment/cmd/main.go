package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"hexabank/internal/observability/metrics"
	"hexabank/internal/observability/tracing"
	fraudclient "hexabank/services/payment/adapters/fraud-client"
	"hexabank/services/payment/adapters/http"
	"hexabank/services/payment/adapters/kafka"
	"hexabank/services/payment/adapters/postgres"
	"hexabank/services/payment/domain/service"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func main() {
	// POSTGRES
	db, err := sqlx.Connect("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatalln(err)
	}

	// TRACING
	tracerName := "payment-service"
	shutdown := tracing.NewTracer(tracerName)
	defer func() {
		if err := shutdown(context.Background()); err != nil {
			log.Fatalf("Failed to shut down tracer: %v", err)
		}
	}()

	// METRICS
	metrics.NewtMetricsEndpoint()

	// FRAUD CLIENT
	fraudClient, err := fraudclient.NewFraudClient(os.Getenv("FRAUD_SERVICE_ADDRESS"))
	if err != nil {
		log.Fatalln(err)
	}

	// KAFKA PRODUCER
	kafkaBrokers := []string{os.Getenv("KAFKA_BROKER_ADDRESS")}
	kafkaTopic := "notifications"
	notificationProducer, err := kafka.NewNotificationProducer(kafkaBrokers, kafkaTopic)
	if err != nil {
		log.Fatalf("Failed to create Kafka producer: %v", err)
	}

	// PAYMENT SERVICE
	paymentRepository := postgres.NewPaymentRepo(db)
	paymentService := service.NewPaymentService(paymentRepository, fraudClient, notificationProducer)
	paymentHandler := http.NewPaymentHTTP(paymentService)

	r := gin.Default()
	paymentHandler.RegisterRoutes(r)

	fmt.Println("Payment service is running on port 8080")
	log.Fatal(r.Run(":8080"))
}
