package services

import (
	"context"
	"fmt"
	"time"

	"github.com/kishanknows/order-service/internal/client"
	"github.com/kishanknows/order-service/internal/config"
	"github.com/kishanknows/order-service/internal/database"
	"github.com/kishanknows/order-service/internal/dto"
	"github.com/kishanknows/order-service/internal/errors"
)

type OrderService struct{
	ProductClient *client.ProductClient
}

func NewOrderService() *OrderService {
	return &OrderService{
		ProductClient: client.NewProductClient(config.Conf.ProductClient.BaseURL),
	}
}

func (s *OrderService) CreateOrder(
	ctx context.Context, 
	userID int,
	req *dto.CreateOrderRequest) (*dto.CreateOrderResponse, *errors.AppError) {
		
	products := []dto.ProductData{}
	var total float64

	for _, item := range req.Items {
		data, err := s.ProductClient.GetProduct(ctx, item.ProductID)
		if err != nil {
			return nil, errors.ErrInternalServer.WithError(err)
		}
		if data.Stock < item.Quantity {
			fmt.Printf(`%s item is not in the stock`, item.ProductID)
			return nil, errors.ErrInternalServer
		}
		products = append(products, *data)
	}

	for i, item := range req.Items {
		err := s.ProductClient.UpdateProduct(ctx, item.ProductID, products[i].Stock - item.Quantity)
		if err != nil {
			return nil, errors.ErrInternalServer.WithError(err)
		}
		total += float64(products[i].Price * item.Quantity)
	}

	tx, err := database.DB.Begin()

	if err != nil {
		return nil, errors.ErrInternalServer.WithError(err)
	}

	defer tx.Rollback()

	var orderID int
	err = tx.QueryRow(`
		INSERT INTO orders (user_id, total, created_at)
		VALUES ($1, $2, $3)
		RETURNING id
	`,userID, total, time.Now()).Scan(&orderID)

	if err != nil {
		return nil, errors.ErrInternalServer.WithError(err)
	}

	for i, item := range req.Items {
		_, err := tx.Exec(`
			INSERT INTO order_items (order_id, product_id, price, quantity)
			VALUES ($1, $2, $3, $4)
		`, orderID, item.ProductID, products[i].Price, item.Quantity)

		if err != nil {
			return nil, errors.ErrInternalServer.WithError(err)
		}
	}

	tx.Commit()

	return &dto.CreateOrderResponse{
		Total: total,
		OrderID: orderID,
	}, nil
}