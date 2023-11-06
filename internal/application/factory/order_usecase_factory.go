package factory

import (
	"database/sql"
	"os"
	"payment-service/internal/application/usecase"
	paymentgateway "payment-service/internal/infra/payment_gateway"
	"payment-service/internal/infra/repository"
)

func GetOrderUsecaseFactory(db *sql.DB) *usecase.GetOrderUseCase {
	orderRepository := repository.NewOrderRepository(db)
	getOrderUseCase := usecase.NewGetOrderUseCase(orderRepository)
	return getOrderUseCase
}

func PaymentUseCaseUsecaseFactory(db *sql.DB) *usecase.PaymentUseCase {
	orderRepository := repository.NewOrderRepository(db)
	transactionRepository := repository.NewTransactionRepository(db)
	customerRepository := repository.NewCustomerRepository(db)
	gatewayPayment := paymentgateway.StripeCard{
		Secret: os.Getenv("STRIPE_SK"),
		Key: os.Getenv("STRIPE_PK"),
		Currency: "brl",
	}
	updateOrderStatusUseCase := usecase.NewPaymentUseCase(orderRepository,transactionRepository,customerRepository,&gatewayPayment)
	return updateOrderStatusUseCase
}

