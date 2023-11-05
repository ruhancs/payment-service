package usecase

import (
	"context"
	"fmt"
	"payment-service/internal/application/dto"
	"payment-service/internal/domain/entity"
	"payment-service/internal/domain/gateway"
)

type CreateOrderUseCase struct {
	OrderRepository gateway.OrderRepository
}

func NewCreateOrderUseCase(repository gateway.OrderRepository) *CreateOrderUseCase {
	return &CreateOrderUseCase{
		OrderRepository: repository,
	}
}

func(usecase *CreateOrderUseCase) Execute(ctx context.Context,input dto.CreateOrderInputDto) (dto.CreateOrderOutputDto,error) {
	order,err := entity.NewOrder(input.ID,input.Plan,input.FirstName,input.LastName,input.Email,input.Amount)
	if err != nil {
		fmt.Println("error to create order entity")
		return dto.CreateOrderOutputDto{},err
	}

	fmt.Println(order)
	err = usecase.OrderRepository.Create(ctx,order)
	if err != nil {
		fmt.Println("error to insert order on db")
		fmt.Println(err)
		return dto.CreateOrderOutputDto{},err
	}

	output := dto.CreateOrderOutputDto{
		ID: order.ID,
		Amount: order.Amount,
		Plan: order.Plan,
		FirstName: order.FirstName,
		LastName: order.LastName,
		Email: order.Email,
		Status: order.Status,
		CreatedAt: order.CreatedAt,
	}

	return output,nil
}