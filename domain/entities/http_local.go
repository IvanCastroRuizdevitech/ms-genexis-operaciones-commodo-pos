package entities

type Response struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
	Error   string      `json:"error,omitempty"`
}

func NewSuccessResponse(message string, data interface{}) Response {
	return Response{
		Success: true,
		Message: message,
		Data:    data,
	}
}
func NewResponseDataBase(message string, data interface{}, success bool) Response {
	return Response{
		Success: success,
		Message: message,
		Data:    data,
	}
}

func NewErrorResponse(message string, err error) Response {
	return Response{
		Success: false,
		Message: message,
		Error:   err.Error(),
	}
}
