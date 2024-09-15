package ports

import (
	"context"

	"github.com/andrei-kozel/microservices-demo/order/internal/application/core/domain"
)

type APIPort interface {
	PlaceOrder(ctx context.Context, order domain.Order) (domain.Order, error)
	GetOrder(ctx context.Context, orderID int64) (domain.Order, error)
}
