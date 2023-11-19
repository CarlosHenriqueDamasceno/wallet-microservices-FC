package create_transaction

import (
	"context"
	"fmt"

	"github.com/CarlosHenriqueDamasceno/wallet-ms-fc/internal/entity"
	"github.com/CarlosHenriqueDamasceno/wallet-ms-fc/internal/gateway"
	"github.com/CarlosHenriqueDamasceno/wallet-ms-fc/pkg/events"
	"github.com/CarlosHenriqueDamasceno/wallet-ms-fc/pkg/uow"
)

type CreateTransactionInputDto struct {
	AccountIDFrom string  `json:"account_id_from"`
	AccountIDTo   string  `json:"account_id_to"`
	Amount        float64 `json:"amount"`
}

type CreateTransactionOutputDto struct {
	ID            string  `json:"id"`
	AccountIDFrom string  `json:"account_id_from"`
	AccountIDTo   string  `json:"account_id_to"`
	Amount        float64 `json:"amount"`
}

type CreateTransactionUseCase struct {
	Uow                uow.UowInterface
	EventDispatcher    events.EventDispatcherInterface
	TransactionCreated events.EventInterface
}

func NewCreateTransactionUseCase(
	uow uow.UowInterface,
	eventDispatcher events.EventDispatcherInterface,
	transactionCreated events.EventInterface,
) *CreateTransactionUseCase {
	return &CreateTransactionUseCase{
		Uow:                uow,
		EventDispatcher:    eventDispatcher,
		TransactionCreated: transactionCreated,
	}
}

func (u *CreateTransactionUseCase) Execute(ctx context.Context, input CreateTransactionInputDto) (*CreateTransactionOutputDto, error) {
	output := &CreateTransactionOutputDto{}
	err := u.Uow.Do(ctx, func(_ *uow.Uow) error {
		accountRepository := u.getAccountRepository(ctx)
		transactionRepository := u.getTransactionRepository(ctx)
		accountFrom, err := accountRepository.Get(input.AccountIDFrom)
		if err != nil {
			return err
		}
		accountTo, err := accountRepository.Get(input.AccountIDTo)
		if err != nil {
			return err
		}
		transaction, err := entity.NewTransaction(accountFrom, accountTo, input.Amount)
		if err != nil {
			return err
		}
		fmt.Println(accountFrom.Balance)
		err = accountRepository.UpdateBalance(accountFrom)
		if err != nil {
			return err
		}
		err = accountRepository.UpdateBalance(accountTo)
		if err != nil {
			return err
		}
		err = transactionRepository.Save(transaction)
		if err != nil {
			return err
		}
		output.ID = transaction.ID
		output.AccountIDFrom = input.AccountIDFrom
		output.AccountIDTo = input.AccountIDTo
		output.Amount = input.Amount
		return nil
	})
	if err != nil {
		return nil, err
	}
	u.TransactionCreated.SetPayload(output)
	u.EventDispatcher.Dispatch(u.TransactionCreated)
	return output, nil
}

func (uc *CreateTransactionUseCase) getAccountRepository(ctx context.Context) gateway.AccountGateway {
	repo, err := uc.Uow.GetRepository(ctx, "AccountDB")
	if err != nil {
		panic(err)
	}
	return repo.(gateway.AccountGateway)
}

func (uc *CreateTransactionUseCase) getTransactionRepository(ctx context.Context) gateway.TransactionGateway {
	repo, err := uc.Uow.GetRepository(ctx, "TransactionDB")
	if err != nil {
		panic(err)
	}
	return repo.(gateway.TransactionGateway)
}
