package gateway

import "github.com/CarlosHenriqueDamasceno/wallet-ms-fc/internal/entity"

type TransactionGateway interface {
	Save(*entity.Transaction) error
}
