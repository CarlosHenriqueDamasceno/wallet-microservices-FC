package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateNewClient(t *testing.T) {
	client, err := NewClient("Carlos", "carlos@teste.com")
	assert.Nil(t, err)
	assert.NotNil(t, client)
	assert.Equal(t, "Carlos", client.Name)
	assert.Equal(t, "carlos@teste.com", client.Email)
}

func TestCreateNewClientWhenNameIsEmpty(t *testing.T) {
	client, err := NewClient("", "carlos@teste.com")
	assert.NotNil(t, err)
	assert.Nil(t, client)
	assert.Equal(t, "name is required", err.Error())
}

func TestCreateNewClientWhenEmailIsEmpty(t *testing.T) {
	client, err := NewClient("Carlos", "")
	assert.NotNil(t, err)
	assert.Nil(t, client)
	assert.Equal(t, "email is required", err.Error())
}

func TestUpdateClient(t *testing.T) {
	client, _ := NewClient("Carlos", "carlos@teste.com")
	err := client.Update("Jhon Doe", "jhondoe@teste.com")
	assert.Nil(t, err)
	assert.Equal(t, "Jhon Doe", client.Name)
	assert.Equal(t, "jhondoe@teste.com", client.Email)
}

func TestUpdateClientWhenNameIsEmpty(t *testing.T) {
	client, _ := NewClient("Carlos", "carlos@teste.com")
	err := client.Update("", "jhondoe@teste.com")
	assert.Error(t, err, "name is required")
}

func TestAddAccountToClient(t *testing.T) {
	client, _ := NewClient("Carlos", "carlos@teste.com")
	account := NewAccount(client)
	err := client.AddAccount(account)
	assert.Nil(t, err)
	assert.Equal(t, 1, len(client.Accounts))

}
