package main

import (
	"database/sql"
	"os"
	"payment-service/internal/application/factory"
	"payment-service/internal/infra/web"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	conn,err := sql.Open(os.Getenv("DB_DRIVER"), os.Getenv("DSN"))
	if err != nil{
		panic(err)
	}
	defer conn.Close()

	getOrderUseCase := factory.GetOrderUsecaseFactory(conn)
	paymentUseCase := factory.PaymentUseCaseUsecaseFactory(conn)
	app := web.NewApplication(getOrderUseCase,paymentUseCase)

	app.Server()
}