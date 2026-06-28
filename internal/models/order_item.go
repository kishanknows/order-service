package models

type OrderItem struct {
	ID int `json:"id"`
	OrderID int `json:"order_id"`
	ProductID string `json:"product_id"`
	Price float64 `json:"price"`
	Quantity int `json:"quantity"`
}