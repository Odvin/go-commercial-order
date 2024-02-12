package ports

import "github.com/Odvin/go-commercial-order/internal/application/core/domain"

type APIPorts interface {
	PlaceOrder(order domain.Order) (domain.Order, error)
}
