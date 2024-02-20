package port

import "github.com/Odvin/go-commercial-order/application/domain"

type Payment interface {
	Charge(*domain.Order) error
}
