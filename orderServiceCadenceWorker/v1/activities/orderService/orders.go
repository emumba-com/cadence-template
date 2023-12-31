package orderService

import (
	"context"
	"fmt"
)

func ShipOrder(_ context.Context, orderID string) error {
	fmt.Println("Order shipped successfully", orderID)

	return nil
}

func UpdateOrderStatus(_ context.Context, orderID string, status string) error {
	fmt.Println("order status updated to ", status)

	return nil
}
