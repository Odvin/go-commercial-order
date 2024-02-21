package db

import (
	"github.com/Odvin/go-commercial-order/application/domain"
)

func (a Adapter) Get(id string) (domain.Order, error) {
	var orderEntity Order
	var orderItems []domain.OrderItem

	res := a.db.First(&orderEntity, id)

	for _, orderItem := range orderEntity.OrderItems {
		orderItems = append(orderItems, domain.OrderItem{
			ProductCode: orderItem.ProductCode,
			UnitPrice:   orderItem.UnitPrice,
			Quantity:    orderItem.Quantity,
		})
	}

	order := domain.Order{
		ID:         int64(orderEntity.ID),
		CustomerID: orderEntity.CustomerID,
		Status:     orderEntity.Status,
		OrderItems: orderItems,
		CreatedAt:  orderEntity.CreatedAt.UnixNano(),
	}

	return order, res.Error
}

func (a Adapter) Save(order *domain.Order) error {
	var orderItems []OrderItem

	for _, orderItem := range order.OrderItems {
		orderItems = append(orderItems, OrderItem{
			ProductCode: orderItem.ProductCode,
			UnitPrice:   orderItem.UnitPrice,
			Quantity:    orderItem.Quantity,
		})
	}

	orderModel := Order{
		CustomerID: order.CustomerID,
		Status:     order.Status,
		OrderItems: orderItems,
	}

	res := a.db.Create(&orderModel)

	if res.Error == nil {
		order.ID = int64(orderModel.ID)
	}

	return res.Error
}
