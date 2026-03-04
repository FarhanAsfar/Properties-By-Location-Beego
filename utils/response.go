package utils

// global response structure
type APIResponse[T any] struct {
	Success bool   `json:"success"`
	Data    *T     `json:"data,omitempty"`
	Error   string `json:"error,omitempty"`
}

// successful api response
func SuccessResponse[T any](data T) APIResponse[T] {
	return APIResponse[T]{
		Success: true,
		Data:    &data,
	}
}

// error api response
func ErrorResponse(message string) APIResponse[any] {
	return APIResponse[any]{
		Success: false,
		Error:   message,
	}
}
