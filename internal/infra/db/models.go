// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.21.0

package db

import (
	"database/sql"
	"time"
)

type Order struct {
	ID            string
	Amount        int32
	Plan          string
	CustomerID    string
	FirstName     string
	LastName      sql.NullString
	Email         string
	Status        string
	TransactionID sql.NullString
	CreatedAt     time.Time
}

type Transaction struct {
	ID                string
	Amount            int32
	Currency          string
	PaymentIntent     string
	PaymentMethod     string
	ExpireMonth       string
	ExpireYear        string
	TransactionStatus sql.NullString
	CreatedAt         time.Time
	UpdatedAt         sql.NullTime
}