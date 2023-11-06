package usecase

import (
	"context"
	"payment-service/internal/application/dto"
	"payment-service/internal/domain/entity"
	"payment-service/internal/domain/gateway"
	paymentgateway "payment-service/internal/infra/payment_gateway"
)

type PaymentUseCase struct {
	OrderRepository       gateway.OrderRepository
	TransactionRepository gateway.TransactionRepositoryInterface
	CustomerRepository gateway.CustomerRepositoryInterface
	PaymentGateway        paymentgateway.PaymentGatewayInterface
}

func NewPaymentUseCase(
	orderRepository gateway.OrderRepository,
	transactionRepository gateway.TransactionRepositoryInterface,
	customerRepository gateway.CustomerRepositoryInterface, 
	paymentGateway paymentgateway.PaymentGatewayInterface,
) *PaymentUseCase {
	return &PaymentUseCase{
		OrderRepository: orderRepository,
		TransactionRepository: transactionRepository,
		CustomerRepository: customerRepository,
		PaymentGateway:  paymentGateway,
	}
}

func (usecase *PaymentUseCase) Execute(ctx context.Context, input dto.PaymentInputDto) (dto.UpdateOrderStatusOutputDto, string, error) {
	check := true
	var statusMSG = ""
	var subsID string
	var customerID = ""

	customer,err := usecase.CustomerRepository.GetByEmail(ctx,input.Email)
	if err != nil {
		customerCreatedID, msgErr, err := usecase.PaymentGateway.CreateCustomer(input.PaymentMethod, input.Email)
		if err != nil {
			check = false
			statusMSG = msgErr
		}
		customerID = customerCreatedID
	}

	if customerID == "" {
		customerID = customer.ID
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
		_,err = usecase.TransactionRepository.Create(ctx,transactionEntity, statusMSG)
	}
	
	orderEntity,err := entity.NewOrder(input.Plan,input.FirstName,input.LastName,input.Email, statusMSG, transactionEntity.ID,input.Amount)
	if err != nil {
		return dto.UpdateOrderStatusOutputDto{},"error to create order",err
	}
	err = usecase.OrderRepository.Create(ctx,orderEntity)
	if err != nil {
		err = usecase.OrderRepository.Create(ctx,orderEntity)
	}

	output := dto.UpdateOrderStatusOutputDto{
		ID:            orderEntity.ID,
		Amount:        orderEntity.Amount,
		Plan:          orderEntity.Plan,
		FirstName:     orderEntity.FirstName,
		LastName:      orderEntity.LastName,
		Email:         orderEntity.Email,
		Status:        orderEntity.Status,
		TransactionID: orderEntity.TransactionID,
		CreatedAt:     orderEntity.CreatedAt,
	}

	return output, statusMSG, nil
}
