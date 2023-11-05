package repository

import (
	"context"
	"database/sql"
	"payment-service/internal/domain/entity"
	"payment-service/internal/infra/db"
)

type OrderRepository struct {
	DB      *sql.DB
	Queries *db.Queries
}

func NewOrderRepository(database *sql.DB) *OrderRepository {
	return &OrderRepository{
		DB: database,
		Queries: db.New(database),
	}
}

func (repo *OrderRepository) Create(ctx context.Context,order *entity.Order) error {
	err := repo.Queries.CreateOrder(ctx,db.CreateOrderParams{
		ID: order.ID,
		Amount: int32(order.Amount),
		Plan: order.Plan,
		FirstName: order.FirstName,
		LastName: sql.NullString{String: order.LastName, Valid: true},
		Email: order.Email,
		CustomerID: order.CustomerID,
		Status: order.Status,
		CreatedAt: order.CreatedAt,
	})

	if err != nil {
		return err
	}

	return nil
}

func(repo *OrderRepository) Get(ctx context.Context,id string) (*entity.Order, error) {
	orderModel,err := repo.Queries.GetOrders(ctx,id)
	if err != nil {
		return nil,err
	}

	orderEntity := &entity.Order{
		ID: orderModel.ID,
		Amount: int(orderModel.Amount),
		Plan: orderModel.Plan,
		CustomerID: orderModel.CustomerID,
		FirstName: orderModel.FirstName,
		LastName: orderModel.LastName.String,
		Email: orderModel.Email,
		Status: orderModel.Status,
		TransactionID: orderModel.TransactionID.String,
		CreatedAt: orderModel.CreatedAt,
	}

	return orderEntity,nil
}

func(repo *OrderRepository) UpdateStatus(ctx context.Context,orderID, status, transactionID, customerID string) (*entity.Order, error) {
	orderModel,err := repo.UpdateStatus(ctx,orderID,status,transactionID,customerID)
	if err != nil {
		return nil,err
	}

	orderEntity := &entity.Order{
		ID: orderModel.ID,
		Amount: int(orderModel.Amount),
		Plan: orderModel.Plan,
		CustomerID: orderModel.CustomerID,
		FirstName: orderModel.FirstName,
		LastName: orderModel.LastName,
		Email: orderModel.Email,
		Status: orderModel.Status,
		TransactionID: orderModel.TransactionID,
		CreatedAt: orderModel.CreatedAt,
	}

	return orderEntity,nil
}
