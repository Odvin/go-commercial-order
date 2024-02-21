package payment

import (
	"context"

	"github.com/Odvin/go-commercial-order/application/domain"
	"github.com/Odvin/go-commercial-proto/golang/payment"
)

func (a *Adapter) Charge(order *domain.Order) error {
	_, err := a.payment.Create(context.Background(), &payment.CreatePaymentRequest{
		UserId:     order.CustomerID,
		OrderId:    order.ID,
		TotalPrice: order.TotalPrice(),
	})

	return err
}
