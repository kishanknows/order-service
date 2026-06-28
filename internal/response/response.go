package response

import (
	"github.com/gin-gonic/gin"
	"github.com/kishanknows/order-service/internal/errors"
)

type APIResponse struct {
	Success bool `json:"success"`
	Message string `json:"message,omitempty"`
	Data any `json:"data,omitempty"`
	Error error `json:"error,omitempty"`
}

func Success(ctx *gin.Context, code int, message string, data any) {
	ctx.JSON(code, APIResponse{
		Success: true,
		Message: message,
		Data: data,
	})
}

func Failure(ctx *gin.Context, err *errors.AppError) {
	ctx.JSON(err.Code, APIResponse{
		Success: false,
		Message: err.Message,
		Error: err.Err,
	})
}