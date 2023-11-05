package gateway

import (
	"context"
	"payment-service/internal/domain/entity"
)


type TransactionRepositoryInterface interface {
	Create(ctx context.Context,transaction *entity.Transaction, status string) (*entity.Transaction,error)
}