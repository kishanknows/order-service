package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kishanknows/order-service/internal/dto"
	"github.com/kishanknows/order-service/internal/errors"
	"github.com/kishanknows/order-service/internal/response"
	"github.com/kishanknows/order-service/internal/services"
)

type OrderHandler struct {
	OrderService *services.OrderService
}

func NewOrderHandler() *OrderHandler {
	return &OrderHandler{
		OrderService: services.NewOrderService(),
	}
}

func (h *OrderHandler) CreateOrder(ctx *gin.Context) {
	var orderRequest dto.CreateOrderRequest
	if err := ctx.ShouldBindJSON(&orderRequest); err != nil {
		response.Failure(ctx, errors.ErrInternalServer.WithError(err))
		return
	}

	userID, ok := ctx.Value("user_id").(int)
	if !ok {
		response.Failure(ctx, errors.ErrUnauthorized)
		return
	}

	res, err := h.OrderService.CreateOrder(ctx.Request.Context(), userID, &orderRequest)

	if err != nil {
		response.Failure(ctx, err)
		return
	}

	response.Success(ctx, http.StatusOK, "order placed successfully", *res)
}