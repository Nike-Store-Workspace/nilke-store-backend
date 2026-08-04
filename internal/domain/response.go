package domain

type BaseResponse[T any] struct {
	Success   bool   `json:"success"`
	Message   string `json:"message,omitempty"`
	ErrorCode string `json:"error_code,omitempty"`
	Data      *T     `json:"data,omitempty"`
}

// PaginatedData برای پاسخ‌های لیستی همراه با تعداد صفحات
type PaginatedData[T any] struct {
	Items      []T `json:"items"`
	Page       int `json:"page"`
	Limit      int `json:"limit"`
	TotalItems int `json:"total_items,omitempty"`
	TotalPages int `json:"total_pages,omitempty"`
}

func SuccessResponse[T any](data *T, message string) BaseResponse[T] {
	return BaseResponse[T]{
		Success: true,
		Message: message,
		Data:    data,
	}
}

func ErrorResponse(errorCode string, message string) BaseResponse[any] {
	return BaseResponse[any]{
		Success:   false,
		ErrorCode: errorCode,
		Message:   message,
		Data:      nil,
	}
}
