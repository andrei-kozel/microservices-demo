package ports

import "github.com/andrei-kozel/microservices-demo/internal/application/core/domain"

type APIPort interface {
	PlaceOrder(order domain.Order) (domain.Order, error)
}
