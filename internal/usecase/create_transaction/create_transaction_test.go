package create_transaction

import (
	"context"
	"testing"

	"github.com/CarlosHenriqueDamasceno/wallet-ms-fc/internal/entity"
	"github.com/CarlosHenriqueDamasceno/wallet-ms-fc/internal/event"
	"github.com/CarlosHenriqueDamasceno/wallet-ms-fc/internal/usecase/mocks"
	"github.com/CarlosHenriqueDamasceno/wallet-ms-fc/pkg/events"
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

func (m *AccountGatewayMock) UpdateBalance(account *entity.Account) error {
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

	mockUow := &mocks.UowMock{}
	mockUow.On("Do", mock.Anything, mock.Anything).Return(nil)

	input := CreateTransactionInputDto{
		AccountIDFrom: accountFrom.ID,
		AccountIDTo:   accountTo.ID,
		Amount:        100,
	}

	dispatcher := events.NewEventDispatcher()
	event := event.NewTransactionCreated()
	ctx := context.Background()

	useCase := NewCreateTransactionUseCase(mockUow, dispatcher, event)
	output, err := useCase.Execute(ctx, input)
	assert.Nil(t, err)
	assert.NotNil(t, output)
	mockUow.AssertExpectations(t)
	mockUow.AssertNumberOfCalls(t, "Do", 1)
}
