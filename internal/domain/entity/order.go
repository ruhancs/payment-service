package entity

import (
	"time"

	"github.com/asaskevich/govalidator"
	uuid "github.com/satori/go.uuid"
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

func NewOrder(plan, firstName, lastName, Email, status, transactionID string, amount int) (*Order, error) {
	order := &Order{
		ID:            uuid.NewV4().String(),
		Amount:        amount,
		Plan:          plan,
		FirstName:     firstName,
		LastName:      lastName,
		Email:         Email,
		Status:        status,
		TransactionID: transactionID,
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
