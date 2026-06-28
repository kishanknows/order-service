package dto

type CreateOrderRequest struct {
	Items []OrderItemRequest `json:"items"`
}

type OrderItemRequest struct {
	ProductID string `json:"product_id"`
	Quantity int `json:"quantity"`
}