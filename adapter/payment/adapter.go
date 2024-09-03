package payment

import (
	"github.com/Odvin/go-commercial-proto/golang/payment"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Adapter struct {
	payment payment.PaymentClient
}

func InitAdapter(paymentServiceUrl string) (*Adapter, error) {
	var opts []grpc.DialOption

	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))

	conn, err := grpc.Dial(paymentServiceUrl, opts...)
	if err != nil {
		return nil, err
	}
	// defer conn.Close()

	client := payment.NewPaymentClient(conn)
	return &Adapter{payment: client}, nil
}
