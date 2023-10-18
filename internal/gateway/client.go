package gateway

import "github.com/CarlosHenriqueDamasceno/wallet-ms-fc/internal/entity"

type ClientGateway interface {
	Get(id string) (*entity.Client, error)
	Save(*entity.Client) error
}
