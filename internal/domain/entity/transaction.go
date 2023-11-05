package entity

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

type Transaction struct {
	ID            string `json:"id"`
	Amount        int    `json:"amount"`
	Currency      string `json:"currency"`
	ExpireMonth   int    `json:"expire_month"`
	ExpireYear    int    `json:"expire_year"`
	PaymentIntent string `json:"payment_intent"`
	PaymentMethod string `json:"payment_method"`
	//BankReturnCode    string    `json:"bank_return_code"`
	TransactionStatus string    `json:"transaction_status"`
	CreatedAt         time.Time `json:"created_at"`
	UpdatedAt         time.Time `json:"updated_at"`
}

func NewTransaction(currency, paymentIntent, paymentMethod, transactionStatus string, amount, expireYear, expireMonth int) *Transaction {
	return &Transaction{
		ID:            uuid.NewV4().String(),
		Amount:        amount,
		Currency:      currency,
		ExpireMonth:   expireMonth,
		ExpireYear:    expireYear,
		PaymentIntent: paymentIntent,
		PaymentMethod: paymentMethod,
		//BankReturnCode:    bankReturnCode,
		TransactionStatus: transactionStatus,
		CreatedAt:         time.Now(),
		UpdatedAt:         time.Now(),
	}
}
