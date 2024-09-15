package api

import (
	"github.com/andrei-kozel/microservices-demo/order/internal/application/core/domain"
	"github.com/andrei-kozel/microservices-demo/order/internal/ports"
	"golang.org/x/net/context"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Application struct {
	db      ports.DBPort
	payment ports.PaymentPort
}

func NewApplication(db ports.DBPort, payment ports.PaymentPort) *Application {
	return &Application{
		db:      db,
		payment: payment,
	}
}

func (a Application) PlaceOrder(ctx context.Context, order domain.Order) (domain.Order, error) {
	err := a.db.Save(ctx, &order)
	if err != nil {
		return domain.Order{}, err
	}

	paymentErr := a.payment.Charge(ctx, &order)
	if paymentErr != nil {
		st, _ := status.FromError(paymentErr)
		fieldErr := &errdetails.BadRequest_FieldViolation{
			Field:       "payment",
			Description: st.Message(),
		}
		badReq := &errdetails.BadRequest{}
		badReq.FieldViolations = append(badReq.FieldViolations, fieldErr)
		orderStatus := status.New(codes.InvalidArgument, "order creation failed")
		statusithDetails, _ := orderStatus.WithDetails(badReq)
		return domain.Order{}, statusithDetails.Err()
	}

	return order, nil
}

func (a Application) GetOrder(ctx context.Context, orderId int64) (domain.Order, error) {
	order, err := a.db.Get(ctx, orderId)
	if err != nil {
		return domain.Order{}, err
	}

	return order, nil
}
