package errors

import "net/http"

var (
	ErrInternalServer = New(
		http.StatusInternalServerError,
		"internal server error",
		nil,
	)
)