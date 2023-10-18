package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateTransaction(t *testing.T) {
	client, _ := NewClient("Carlos", "carlos@teste.com")
	accountFrom := NewAccount(client)
	accountTo := NewAccount(client)
	accountFrom.Credit(10)
	transaction, err := NewTransaction(accountFrom, accountTo, 5)
	assert.Nil(t, err)
	assert.NotNil(t, transaction)
	assert.Equal(t, 5.0, accountFrom.Balance)
	assert.Equal(t, 5.0, accountTo.Balance)
}

func TestCreateTransactionWithInsufficientFounds(t *testing.T) {
	client, _ := NewClient("Carlos", "carlos@teste.com")
	accountFrom := NewAccount(client)
	accountTo := NewAccount(client)
	accountFrom.Credit(5.0)
	transaction, err := NewTransaction(accountFrom, accountTo, 10)
	assert.Error(t, err, "insufficient founds")
	assert.Nil(t, transaction)
	assert.Equal(t, 5.0, accountFrom.Balance)
	assert.Equal(t, 0.0, accountTo.Balance)
}
