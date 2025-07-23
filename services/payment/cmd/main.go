package main

import (
	"fmt"
	"log"
	"os"

	"hexabank/payment/adapters/http"
	"hexabank/payment/adapters/postgres"
	"hexabank/payment/domain/service"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func main() {
	db, err := sqlx.Connect("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatalln(err)
	}

	paymentRepository := postgres.NewPaymentRepo(db)
	paymentService := service.NewPaymentService(paymentRepository)
	paymentHandler := http.NewPaymentHTTP(paymentService)

	r := gin.Default()
	paymentHandler.RegisterRoutes(r)

	fmt.Println("Payment service is running on port 8080")
	log.Fatal(r.Run(":8080"))
}
