package port

import "github.com/Odvin/go-commercial-order/application/domain"

type DB interface {
	Get(id string) (domain.Order, error)
	Save(*domain.Order) error
}
