package models

import "time"

type Order struct {
	ID int `json:"id"`
	UserID int `json:"user_id"`
	Status OrderStatus `json:"status"`
	Total float64 `json:"total"`
	CreatedAt time.Time `json:"created_at"`
}

type OrderStatus string

const (
	Pending OrderStatus = "pending"
	Confirmed OrderStatus = "confirmed"
	Shipped OrderStatus = "shipped"
	Delivered OrderStatus = "delivered"
	Cancelled OrderStatus = "cancelled"
)