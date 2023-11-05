package repository

import (
	"context"
	"database/sql"
	"payment-service/internal/domain/entity"
	"payment-service/internal/infra/db"
)

type TransactionRepository struct {
	DB *sql.DB
	Queries *db.Queries
}

func NewTransactionRepository(database *sql.DB) *TransactionRepository {
	return &TransactionRepository{
		DB: database,
		Queries: db.New(database),
	}
}

func(repo *TransactionRepository) Create(ctx context.Context,transaction *entity.Transaction, status string) (*entity.Transaction,error) {
	_,err := repo.Queries.CreateTransaction(ctx,db.CreateTransactionParams{
		ID: transaction.ID,
		Amount: int32(transaction.Amount),
		Currency: transaction.Currency,
		PaymentIntent: transaction.PaymentIntent,
		PaymentMethod: transaction.PaymentMethod,
		TransactionStatus: sql.NullString{String: status,Valid: true},
		ExpireMonth: "",// TODO fix db
		ExpireYear: "",
		CreatedAt: transaction.CreatedAt,
		UpdatedAt: sql.NullTime{Time: transaction.UpdatedAt,Valid: true},
	})
	if err != nil {
		return nil,err
	}

	return transaction,nil
}