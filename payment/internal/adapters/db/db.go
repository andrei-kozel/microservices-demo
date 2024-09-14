package db

import (
	"context"
	"fmt"

	"github.com/andrei-kozel/microservices-demo/payment/internal/application/core/domain"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Payment struct {
	gorm.Model
	Status     string  `json:"status"`
	CustomerID int64   `json:"customer_id"`
	OrderID    int64   `json:"order_id"`
	TotalPrice float32 `json:"total_price"`
}

type Adapter struct {
	db *gorm.DB
}

func (a Adapter) Get(ctx context.Context, id string) (domain.Payment, error) {
	var paymentEntity Payment
	res := a.db.WithContext(ctx).First(&paymentEntity, id)
	payment := domain.Payment{
		ID:         int64(paymentEntity.ID),
		CustomerID: paymentEntity.CustomerID,
		OrderID:    paymentEntity.OrderID,
		TotalPrice: paymentEntity.TotalPrice,
		CreatedAt:  paymentEntity.CreatedAt.Unix(),
		Status:     paymentEntity.Status,
	}

	return payment, res.Error
}

func (a Adapter) Save(ctx context.Context, payment *domain.Payment) error {
	orderModel := Payment{
		CustomerID: payment.CustomerID,
		OrderID:    payment.OrderID,
		TotalPrice: payment.TotalPrice,
		Status:     payment.Status,
	}
	res := a.db.WithContext(ctx).Create(&orderModel)
	if res.Error == nil {
		payment.ID = int64(orderModel.ID)
	}
	return res.Error
}

func NewAdapter(dataSourceUrl string) (*Adapter, error) {
	db, err := gorm.Open(mysql.Open(dataSourceUrl), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to open database connection: %w", err)
	}
	err = db.AutoMigrate(&Payment{})
	if err != nil {
		return nil, fmt.Errorf("failed to auto migrate database: %w", err)
	}
	return &Adapter{
		db: db,
	}, nil
}
