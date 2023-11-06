package gateway

import (
	"context"
	"payment-service/internal/domain/entity"
)

type CustomerRepositoryInterface interface {
	Create(ctx context.Context,customer *entity.Customer) error
	GetByEmail (ctx context.Context,email string) (*entity.Customer,error)
	GetByID (ctx context.Context,id string) (*entity.Customer,error)
}