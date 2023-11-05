package gateway

import (
	"context"
	"payment-service/internal/domain/entity"
)

type OrderRepository interface {
	Create(ctx context.Context,order *entity.Order) error
	Get(ctx context.Context,id string) (*entity.Order, error)
	UpdateStatus(ctx context.Context,orderID, status, transactionID, customerID string) (*entity.Order, error)
}
