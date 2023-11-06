package repository

import (
	"context"
	"database/sql"
	"payment-service/internal/domain/entity"
	"payment-service/internal/infra/db"
)

type CustomerRepository struct {
	DB      *sql.DB
	Queries *db.Queries
}

func NewCustomerRepository(database *sql.DB) *CustomerRepository {
	return &CustomerRepository{
		DB: database,
		Queries: db.New(database),
	}
}

func (repo *CustomerRepository) Create(ctx context.Context,customer *entity.Customer) error {
	err := repo.Queries.CreateCustomer(ctx,db.CreateCustomerParams{
		ID: customer.ID,
		FirstName: customer.FirstName,
		LastName: sql.NullString{String: customer.LastName, Valid: true},
		Email: customer.Email,
		IsActive: customer.IsActive,
	})
	if err != nil {
		return err
	}

	return nil
}

func (repo *CustomerRepository) GetByEmail (ctx context.Context,email string) (*entity.Customer,error) {
	customerModel,err := repo.Queries.GetCustomerByEmail(ctx,email)
	if err != nil {
		return nil,err
	}

	customerEntity := entity.Customer{
		ID: customerModel.ID,
		FirstName: customerModel.FirstName,
		LastName: customerModel.LastName.String,
		Email: customerModel.Email,
		IsActive: customerModel.IsActive,
	}

	return &customerEntity,nil
}

func (repo *CustomerRepository) GetByID (ctx context.Context,id string) (*entity.Customer,error) {
	customerModel,err := repo.Queries.GetCustomerById(ctx,id)
	if err != nil {
		return nil,err
	}

	customerEntity := entity.Customer{
		ID: customerModel.ID,
		FirstName: customerModel.FirstName,
		LastName: customerModel.LastName.String,
		Email: customerModel.Email,
		IsActive: customerModel.IsActive,
	}

	return &customerEntity,nil
}
