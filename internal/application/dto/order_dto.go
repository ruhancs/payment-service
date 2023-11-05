package dto

import "time"

type CreateOrderInputDto struct {
	ID        string `json:"id"`
	Amount    int    `json:"amount"`
	Plan      string `json:"plan"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
}

type CreateOrderOutputDto struct {
	ID        string    `json:"id"`
	Amount    int       `json:"amount"`
	Plan      string    `json:"plan"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Email     string    `json:"email"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
}

type GetOrderOutputDto struct {
	ID            string    `json:"id"`
	Amount        int       `json:"amount"`
	Plan          string    `json:"plan"`
	FirstName     string    `json:"first_name"`
	LastName      string    `json:"last_name"`
	Email         string    `json:"email"`
	Status        string    `json:"status"`
	TransactionID string    `json:"transaction"`
	CreatedAt     time.Time `json:"created_at"`
}

type PaymentInputDto struct {
	Currency      string `json:"currency"`
	Amount        int    `json:"amount"`
	PaymentMethod string `json:"payment_method"`
	Email         string `json:"email"`
	CardBrand     string `json:"card_brand"`
	ExpiryMonth   int    `json:"exp_month"`
	ExpiryYear    int    `json:"exp_year"`
	Plan          string `json:"plan"`
	ProductID     string `json:"product_id"`
	FirstName     string `json:"first_name"`
	LastName      string `json:"last_name"`
}

type UpdateOrderStatusInputDto struct {
	OrderID       string
	Status        string
	CustomerID    string
	TransactionID string
}

type UpdateOrderStatusOutputDto struct {
	ID            string    `json:"id"`
	Amount        int       `json:"amount"`
	Plan          string    `json:"plan"`
	FirstName     string    `json:"first_name"`
	LastName      string    `json:"last_name"`
	Email         string    `json:"email"`
	Status        string    `json:"status"`
	TransactionID string    `json:"transaction"`
	CreatedAt     time.Time `json:"created_at"`
}
