package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/kishanknows/order-service/internal/config"
	"github.com/kishanknows/order-service/internal/handlers"
	middleware "github.com/kishanknows/order-service/internal/middlewares"
)

func RegisterOrderRoutes(r *gin.Engine) {
	handler := handlers.NewOrderHandler()

	r.Use(middleware.AuthMiddleware([]byte(config.Conf.JWTSecret)))
	r.POST("/order", handler.CreateOrder)
}