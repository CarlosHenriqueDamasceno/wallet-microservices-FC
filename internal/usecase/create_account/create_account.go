package createaccount

import (
	"github.com/CarlosHenriqueDamasceno/wallet-ms-fc/internal/entity"
	"github.com/CarlosHenriqueDamasceno/wallet-ms-fc/internal/gateway"
)

type CreateAccountInputDto struct {
	ClientID string
}

type CreateAccountOutputDto struct {
	ID string
}

type CreateAccountUseCase struct {
	AccountGateway gateway.AccountGateway
	ClientGateway  gateway.ClientGateway
}

func NewCreateAccountUseCase(
	accountGateway gateway.AccountGateway,
	clientGateway gateway.ClientGateway,
) *CreateAccountUseCase {
	return &CreateAccountUseCase{
		AccountGateway: accountGateway,
		ClientGateway:  clientGateway,
	}
}

func (u *CreateAccountUseCase) Execute(input CreateAccountInputDto) (*CreateAccountOutputDto, error) {
	client, err := u.ClientGateway.Get(input.ClientID)
	if err != nil {
		return nil, err
	}
	account := entity.NewAccount(client)
	err = u.AccountGateway.Save(account)
	if err != nil {
		return nil, err
	}
	return &CreateAccountOutputDto{
		ID: account.ID,
	}, nil
}
