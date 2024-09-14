package ports

import "github.com/andrei-kozel/microservices-demo/internal/application/core/domain"

type DBPort interface {
	Get(id string) (domain.Order, error)
	Save(*domain.Order) error
}
