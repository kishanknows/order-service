package dto

type CreateOrderResponse struct {
	Total float64 `json:"total"`
	OrderID int `json:"order_id"`
}