package models

type OrderDetails struct {
	ID              string `json:"id"`
	UserID          string `json:"userId"`
	ShippingAddress string `json:"shippingAddress"`
	Status          string `json:"status"`
	OrderDate       string `json:"orderDate"`
	ReceiptID       string `json:"receiptId"`
}
