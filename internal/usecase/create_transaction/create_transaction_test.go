package createtransaction

import (
	"testing"

	"github.com/CarlosHenriqueDamasceno/wallet-ms-fc/internal/entity"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type AccountGatewayMock struct {
	mock.Mock
}

func (m *AccountGatewayMock) Get(id string) (*entity.Account, error) {
	args := m.Called(id)
	return args.Get(0).(*entity.Account), args.Error(1)
}

func (m *AccountGatewayMock) Save(account *entity.Account) error {
	args := m.Called(account)
	return args.Error(0)
}

type TransactionGatewayMock struct {
	mock.Mock
}

func (m *TransactionGatewayMock) Save(transaction *entity.Transaction) error {
	args := m.Called(transaction)
	return args.Error(0)
}

func TestCreateTransactionUseCase_Execute(t *testing.T) {
	client, _ := entity.NewClient("Carlos", "carlos@teste.com")
	accountFrom := entity.NewAccount(client)
	accountFrom.Credit(110.0)
	accountTo := entity.NewAccount(client)

	accountMock := &AccountGatewayMock{}
	accountMock.On("Get", accountFrom.ID).Return(accountFrom, nil)
	accountMock.On("Get", accountTo.ID).Return(accountTo, nil)

	transactionMock := &TransactionGatewayMock{}
	transactionMock.On("Save", mock.Anything).Return(nil)

	useCase := NewCreateTransactionUseCase(accountMock, transactionMock)

	input := CreateTransactionInputDto{
		AccountFromId: accountFrom.ID,
		AccountToId:   accountTo.ID,
		Amount:        100,
	}

	output, err := useCase.Execute(input)
	assert.Nil(t, err)
	assert.NotNil(t, output)
	assert.Equal(t, 10.0, accountFrom.Balance)
	assert.Equal(t, 100.0, accountTo.Balance)
	accountMock.AssertExpectations(t)
	accountMock.AssertNumberOfCalls(t, "Get", 2)
	transactionMock.AssertExpectations(t)
	transactionMock.AssertNumberOfCalls(t, "Save", 1)
}
