package create_client

import (
	"time"

	"github.com/CarlosHenriqueDamasceno/wallet-ms-fc/internal/entity"
	"github.com/CarlosHenriqueDamasceno/wallet-ms-fc/internal/gateway"
)

type CreateClientInputDto struct {
	Name  string
	Email string
}

type CreateClientOutputDto struct {
	ID        string
	Name      string
	Email     string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewCreateClientOutputDto(entity entity.Client) *CreateClientOutputDto {
	return &CreateClientOutputDto{
		ID:        entity.ID,
		Name:      entity.Name,
		Email:     entity.Email,
		CreatedAt: entity.CreatedAt,
		UpdatedAt: entity.UpdatedAt,
	}
}

type CreateClientUseCase struct {
	ClientGateway gateway.ClientGateway
}

func NewCreateClientUseCase(clientGateway gateway.ClientGateway) *CreateClientUseCase {
	return &CreateClientUseCase{
		ClientGateway: clientGateway,
	}
}

func (u *CreateClientUseCase) Execute(input CreateClientInputDto) (*CreateClientOutputDto, error) {
	client, err := entity.NewClient(input.Name, input.Email)
	if err != nil {
		return nil, err
	}
	err = u.ClientGateway.Save(client)
	if err != nil {
		return nil, err
	}
	return NewCreateClientOutputDto(*client), nil
}
