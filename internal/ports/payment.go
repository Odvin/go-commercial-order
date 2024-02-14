package ports

import "github.com/Odvin/go-commercial-order/internal/application/core/domain"

type PaymentPort interface {
	Charge(*domain.Order) error
}
