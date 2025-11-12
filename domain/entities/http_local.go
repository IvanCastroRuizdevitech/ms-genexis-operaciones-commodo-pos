package entities

import "time"

type Response[T any] struct {
	Status      int    `json:"status"`
	Message     string `json:"message,omitempty"`
	ProcessDate string `json:"process_date"`
	Data        *T     `json:"data,omitempty"`
	Error       string `json:"error,omitempty"`
}

// Respuesta exitosa
func NewSuccessResponse[T any](status int, message, processDate string, data *T) Response[T] {
	return Response[T]{
		Status:      status,
		Message:     message,
		ProcessDate: processDate,
		Data:        data,
	}
}

// Respuesta de error
func NewErrorResponse[T any](message string, err error) Response[T] {
	return Response[T]{
		Status:      400,
		Message:     message,
		ProcessDate: time.Now().Format("2006-01-02 15:04:05"),
		Error:       err.Error(),
	}
}
