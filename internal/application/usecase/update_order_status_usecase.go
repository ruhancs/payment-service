package usecase

import (
	"context"
	"payment-service/internal/application/dto"
	"payment-service/internal/domain/entity"
	"payment-service/internal/domain/gateway"
	paymentgateway "payment-service/internal/infra/payment_gateway"
)

type UpdateOrderUseCase struct {
	OrderRepository       gateway.OrderRepository
	TransactionRepository gateway.TransactionRepositoryInterface
	PaymentGateway        paymentgateway.PaymentGatewayInterface
}

func NewUpdateOrderUseCase(
	orderRepository gateway.OrderRepository,
	transactionRepository gateway.TransactionRepositoryInterface, 
	paymentGateway paymentgateway.PaymentGatewayInterface,
) *UpdateOrderUseCase {
	return &UpdateOrderUseCase{
		OrderRepository: orderRepository,
		TransactionRepository: transactionRepository,
		PaymentGateway:  paymentGateway,
	}
}

func (usecase *UpdateOrderUseCase) Execute(ctx context.Context,orderID string, input dto.PaymentInputDto) (dto.UpdateOrderStatusOutputDto, string, error) {
	check := true
	var statusMSG = ""
	var subsID string
	customerID, msgErr, err := usecase.PaymentGateway.CreateCustomer(input.PaymentMethod, input.Email)
	if err != nil {
		check = false
		statusMSG = msgErr
	}

	if check {
		subscriptionID, err := usecase.PaymentGateway.SubscribeToPlan(customerID, input.Plan, input.Email, input.CardBrand)
		if err != nil {
			check = false
			statusMSG = "failed to assign plan"
		}
		subsID = subscriptionID
	}

	
	if statusMSG == "" {
		statusMSG = "payment successfuly"
	}
	
	transactionEntity := entity.NewTransaction("BRL",subsID,input.PaymentMethod,statusMSG,input.Amount,input.ExpiryYear,input.ExpiryMonth)
	_,err = usecase.TransactionRepository.Create(ctx,transactionEntity, statusMSG)
	if err != nil {
		return dto.UpdateOrderStatusOutputDto{},"error to create transaction", err
	} 

	order, err := usecase.OrderRepository.UpdateStatus(
		ctx,
		orderID,
		statusMSG,
		transactionEntity.ID,
		customerID,
	)
	if err != nil {
		return dto.UpdateOrderStatusOutputDto{}, "error to update order status", err
	}

	output := dto.UpdateOrderStatusOutputDto{
		ID:            order.ID,
		Amount:        order.Amount,
		Plan:          order.Plan,
		FirstName:     order.FirstName,
		LastName:      order.LastName,
		Email:         order.Email,
		Status:        order.Status,
		TransactionID: order.TransactionID,
		CreatedAt:     order.CreatedAt,
	}

	return output, statusMSG, nil
}
