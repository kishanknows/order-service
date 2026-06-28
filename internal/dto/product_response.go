package dto

type ProductResponse struct {
	Success bool `json:"success,omitempty"`
	Message string `json:"message,omitempty"`
	Data ProductData `json:"data,omitempty"`
}

type ProductData struct {
	Id string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
	Price int `json:"price,omitempty"`
	SellerId int `json:"seller_id,omitempty"`
	Stock int `json:"stock,omitempty"`
	CreatedAt string `json:"created_at,omitempty"`
	UpdatedAt string `json:"updated_at,omitempty"`
}