package usecase

import (
	"context"
	"payment-service/internal/application/dto"
	"payment-service/internal/domain/gateway"
)

type GetOrderUseCase struct {
	OrderRepository gateway.OrderRepository
}

func NewGetOrderUseCase(repository gateway.OrderRepository) *GetOrderUseCase {
	return &GetOrderUseCase{
		OrderRepository: repository,
	}
}

func (usecase *GetOrderUseCase) Execute(ctx context.Context,id string) (dto.GetOrderOutputDto, error) {
	order, err := usecase.OrderRepository.Get(ctx,id)
	if err != nil {
		return dto.GetOrderOutputDto{}, nil
	}

	output := dto.GetOrderOutputDto{
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

	return output,nil
}
