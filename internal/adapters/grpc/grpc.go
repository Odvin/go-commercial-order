package grpc

import (
	"context"

	"github.com/Odvin/go-commercial-order/internal/application/core/domain"
	"github.com/Odvin/go-commercial-proto/golang/order"
)

func (a Adapter) Create(ctx context.Context, request *order.CreateOrderRequest) (*order.CreateOrderResponse, error) {
	var orderItems []domain.OrderItem

	for _, orderItem := range request.Items {
		orderItems = append(orderItems, domain.OrderItem{
			ProductCode: orderItem.ProductCode,
			UnitPrice:   orderItem.UnitPrice,
			Quantity:    orderItem.Quantity,
		})
	}

	newOder := domain.NewOrder(request.UserId, orderItems)

	result, err := a.api.PlaceOrder(newOder)
	if err != nil {
		return nil, err
	}

	return &order.CreateOrderResponse{OrderId: result.ID}, nil
}
