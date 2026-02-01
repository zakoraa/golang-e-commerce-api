package response

import "github.com/gin-gonic/gin"

type SuccessResponse[T any] struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Data    T      `json:"data,omitempty"`
	Meta    any    `json:"meta,omitempty"`
}

func Success[T any](
	c *gin.Context,
	data T,
	message string,
	code int,
	meta ...any,
) {
	res := SuccessResponse[T]{
		Success: true,
		Message: message,
		Data:    data,
	}

	if len(meta) > 0 {
		res.Meta = meta[0]
	}

	c.JSON(code, res)
}

func SuccessNoData(
	c *gin.Context,
	message string,
	code int,
) {
	c.JSON(code, gin.H{
		"success": true,
		"message": message,
	})
}


type ErrorDetail struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
	Code    string `json:"code,omitempty"`
}

type ErrorResponse struct {
	Success bool        `json:"success"`
	Error   ErrorDetail `json:"error"`
}

func Error(
	c *gin.Context,
	message string,
	code int,
	errorCode ...string,
) {
	res := ErrorResponse{
		Success: false,
		Error: ErrorDetail{
			Message: message,
			Status:  code,
		},
	}

	if len(errorCode) > 0 {
		res.Error.Code = errorCode[0]
	}

	c.JSON(code, res)
}
