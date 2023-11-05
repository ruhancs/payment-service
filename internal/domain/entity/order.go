package entity

import (
	"time"

	"github.com/asaskevich/govalidator"
)

type Order struct {
	ID            string    `json:"id"`
	Amount        int       `json:"amount"`
	Plan          string    `json:"plan"`
	CustomerID    string    `json:"customer_id"`
	FirstName     string    `json:"first_name"`
	LastName      string    `json:"last_name"`
	Email         string    `json:"email"`
	Status        string    `json:"status"`
	TransactionID string    `json:"transaction"`
	CreatedAt     time.Time `json:"created_at"`
}

func NewOrder(id,plan, firstName, lastName, Email string, amount int) (*Order, error) {
	order := &Order{
		ID:            id,
		Amount:        amount,
		Plan:          plan,
		FirstName:     firstName,
		LastName:      lastName,
		Email:         Email,
		Status:        "pending",
		TransactionID: "no-transaction",
		CreatedAt:     time.Now(),
	}
	err := order.validate()
	if err != nil {
		return nil, err
	}
	return order, nil
}

func (order *Order) validate() error {
	_, err := govalidator.ValidateStruct(order)
	if err != nil {
		return err
	}
	return nil
}
