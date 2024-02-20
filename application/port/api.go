package port

import "github.com/Odvin/go-commercial-order/application/domain"

type API interface {
	PlaceOrder(order domain.Order) (domain.Order, error)
}
