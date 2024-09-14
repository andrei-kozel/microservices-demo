package domain

import "time"

type Payment struct {
	Status     string  `json:"status"`
	ID         int64   `json:"id"`
	CustomerID int64   `json:"customer_id"`
	OrderID    int64   `json:"order_id"`
	CreatedAt  int64   `json:"created_at"`
	TotalPrice float32 `json:"total_price"`
}

func NewPayment(customerId int64, orderId int64, totalPrice float32) Payment {
	return Payment{
		CustomerID: customerId,
		OrderID:    orderId,
		TotalPrice: totalPrice,
		CreatedAt:  time.Now().Unix(),
		Status:     "Pending",
	}
}
