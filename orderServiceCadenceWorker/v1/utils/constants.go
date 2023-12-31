package utils

const (
	Duration1Hour    = 60
	Duration5Minutes = 5
)

type OrderStatus string

const (
	OrderStatusProcessing OrderStatus = "processing"
	OrderStatusFailed     OrderStatus = "failed"
	OrderStatusCompleted  OrderStatus = "completed"
)
