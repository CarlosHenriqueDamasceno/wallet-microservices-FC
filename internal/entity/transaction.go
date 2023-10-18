package entity

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

type Transaction struct {
	ID          string
	AccountFrom *Account
	AccountTo   *Account
	Amount      float64
	CreatedAt   time.Time
}

func NewTransaction(from *Account, to *Account, amount float64) (*Transaction, error) {
	transaction := &Transaction{
		ID:          uuid.New().String(),
		AccountFrom: from,
		AccountTo:   to,
		Amount:      amount,
		CreatedAt:   time.Now(),
	}
	err := transaction.Validate()
	if err != nil {
		return nil, err
	}
	transaction.commit()
	return transaction, nil
}

func (t *Transaction) commit() {
	t.AccountFrom.Debit(t.Amount)
	t.AccountTo.Credit(t.Amount)
}

func (t *Transaction) Validate() error {
	if t.Amount < 0 {
		return errors.New("amount must be greater then zero")
	}
	if t.AccountFrom.Balance < t.Amount {
		return errors.New("insufficient founds")
	}
	return nil
}
