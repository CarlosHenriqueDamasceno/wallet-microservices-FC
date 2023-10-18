package createtransaction

import (
	"github.com/CarlosHenriqueDamasceno/wallet-ms-fc/internal/entity"
	"github.com/CarlosHenriqueDamasceno/wallet-ms-fc/internal/gateway"
)

type CreateTransactionInputDto struct {
	AccountFromId string
	AccountToId   string
	Amount        float64
}

type CreateTransactionOutputDto struct {
	ID string
}

type CreateTransactionUseCase struct {
	AccountGateway     gateway.AccountGateway
	TransactionGateway gateway.TransactionGateway
}

func NewCreateTransactionUseCase(
	accountGateway gateway.AccountGateway,
	transactionGateway gateway.TransactionGateway,
) *CreateTransactionUseCase {
	return &CreateTransactionUseCase{
		AccountGateway:     accountGateway,
		TransactionGateway: transactionGateway,
	}
}

func (u *CreateTransactionUseCase) Execute(input CreateTransactionInputDto) (*CreateTransactionOutputDto, error) {
	accountFrom, err := u.AccountGateway.Get(input.AccountFromId)
	if err != nil {
		return nil, err
	}
	accountTo, err := u.AccountGateway.Get(input.AccountToId)
	if err != nil {
		return nil, err
	}
	transaction, err := entity.NewTransaction(accountFrom, accountTo, input.Amount)
	if err != nil {
		return nil, err
	}
	err = u.TransactionGateway.Save(transaction)
	if err != nil {
		return nil, err
	}
	return &CreateTransactionOutputDto{
		ID: transaction.ID,
	}, nil
}
