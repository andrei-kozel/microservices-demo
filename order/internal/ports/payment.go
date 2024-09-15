package ports

import (
	"context"

	"github.com/andrei-kozel/microservices-demo/order/internal/application/core/domain"
)

type PaymentPort interface {
	Charge(context.Context, *domain.Order) error
}
