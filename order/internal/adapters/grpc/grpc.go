package grpc

import (
	"context"

	"github.com/andrei-kozel/microservices-demo-proto/golang/order"
	"github.com/andrei-kozel/microservices-demo/order/internal/application/core/domain"
)

func (a Adapter) Create(ctx context.Context, request *order.CreateOrderRequest) (*order.CreateOrderResponse, error) {
	var orderItems []domain.OrderItem
	for _, item := range request.OrderItems {
		orderItems = append(orderItems, domain.OrderItem{
			ProductCode: item.ProductCode,
			UnitPrice:   item.UnitPrice,
			Quantity:    item.Quantity,
		})
	}

	newOrder := domain.NewOrder(request.UserId, orderItems)
	result, err := a.api.PlaceOrder(newOrder)
	if err != nil {
		return nil, err
	}

	return &order.CreateOrderResponse{OrderId: result.ID}, nil
}

func (a Adapter) Get(ctx context.Context, request *order.GetOrderRequest) (*order.GetOrderResponse, error) {
	result, err := a.api.GetOrder(request.OrderId)
	var orderItems []*order.OrderItem
	for _, orderItem := range result.OrderItems {
		orderItems = append(orderItems, &order.OrderItem{
			ProductCode: orderItem.ProductCode,
			UnitPrice:   orderItem.UnitPrice,
			Quantity:    orderItem.Quantity,
		})
	}
	if err != nil {
		return nil, err
	}

	orderItem := &domain.Order{
		ID:         result.ID,
		CustomerID: result.CustomerID,
		CreatedAt:  result.CreatedAt,
		Status:     result.Status,
		OrderItems: result.OrderItems,
	}

	return &order.GetOrderResponse{
		UserId:     orderItem.CustomerID,
		OrderItems: orderItems,
	}, nil
}
