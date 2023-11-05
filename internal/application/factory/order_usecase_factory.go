package factory

import (
	"database/sql"
	"os"
	"payment-service/internal/application/usecase"
	paymentgateway "payment-service/internal/infra/payment_gateway"
	"payment-service/internal/infra/repository"
)


func CreateOrderUsecaseFactory(db *sql.DB) *usecase.CreateOrderUseCase {
	orderRepository := repository.NewOrderRepository(db)
	createOrderUseCase := usecase.NewCreateOrderUseCase(orderRepository)
	return createOrderUseCase
}

func GetOrderUsecaseFactory(db *sql.DB) *usecase.GetOrderUseCase {
	orderRepository := repository.NewOrderRepository(db)
	getOrderUseCase := usecase.NewGetOrderUseCase(orderRepository)
	return getOrderUseCase
}

func UpdateOrderStatusUsecaseFactory(db *sql.DB) *usecase.UpdateOrderUseCase {
	orderRepository := repository.NewOrderRepository(db)
	transactionRepository := repository.NewTransactionRepository(db)
	gatewayPayment := paymentgateway.StripeCard{
		Secret: os.Getenv("STRIPE_SK"),
		Key: os.Getenv("STRIPE_PK"),
		Currency: "brl",
	}
	updateOrderStatusUseCase := usecase.NewUpdateOrderUseCase(orderRepository,transactionRepository,&gatewayPayment)
	return updateOrderStatusUseCase
}

