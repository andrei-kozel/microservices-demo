package payment

import (
	"context"

	"github.com/andrei-kozel/microservices-demo-proto/golang/payment"
	"github.com/andrei-kozel/microservices-demo/order/internal/application/core/domain"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Adapter struct {
	payment payment.PaymentClient
}

func NewAdapter(paymentServiceUrl string) (*Adapter, error) {
	var opts []grpc.DialOption
	opts = append(
		opts,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	conn, err := grpc.NewClient(
		paymentServiceUrl,
		opts...,
	)
	if err != nil {
		return nil, err
	}

	client := payment.NewPaymentClient(conn)

	return &Adapter{payment: client}, nil
}

func (a *Adapter) Charge(ctx context.Context, order *domain.Order) error {
	_, err := a.payment.Create(
		ctx,
		&payment.CreatePaymentRequest{
			UserId:     order.CustomerID,
			OrderId:    order.ID,
			TotalPrice: order.TotalPrice(),
		},
	)

	return err
}
