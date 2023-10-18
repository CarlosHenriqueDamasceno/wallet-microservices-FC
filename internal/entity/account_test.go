package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateAccount(t *testing.T) {
	client, _ := NewClient("Carlos", "carlos@teste.com")
	account := NewAccount(client)
	assert.NotNil(t, account)
	assert.Equal(t, client.ID, account.Client.ID)
}

func TestCreateAccountWithNilClient(t *testing.T) {
	account := NewAccount(nil)
	assert.Nil(t, account)
}

func TestCreditAccount(t *testing.T) {
	client, _ := NewClient("Carlos", "carlos@teste.com")
	account := NewAccount(client)
	account.Credit(10)
	assert.Equal(t, 10.0, account.Balance)
}

func TestDebitAccount(t *testing.T) {
	client, _ := NewClient("Carlos", "carlos@teste.com")
	account := NewAccount(client)
	account.Credit(10)
	assert.Equal(t, 10.0, account.Balance)
	account.Debit(5)
	assert.Equal(t, 5.0, account.Balance)

}
