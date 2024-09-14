package ports

import "github.com/andrei-kozel/microservices-demo/internal/application/core/domain"

type PaymentPort interface {
	Charge(*domain.Order) error
}
