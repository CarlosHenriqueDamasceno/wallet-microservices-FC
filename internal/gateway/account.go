package gateway

import "github.com/CarlosHenriqueDamasceno/wallet-ms-fc/internal/entity"

type AccountGateway interface {
	Get(id string) (*entity.Account, error)
	Save(*entity.Account) error
	UpdateBalance(account *entity.Account) error
}
