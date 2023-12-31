package paymentService

import (
	"context"
	"fmt"
)

func MakePayment(ctx context.Context, orderID string) error {
	fmt.Println("Payment made successfully against orderID:", orderID)

	return nil
}

func RefundPayment(ctx context.Context, orderID string) error {
	fmt.Println("Payment refunded successfully against orderID:", orderID)

	return nil
}
