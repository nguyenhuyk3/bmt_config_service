package responses

import (
	"github.com/gin-gonic/gin"
)

type responseData struct {
	Code    int         `json:"status_code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func SuccessResponse(c *gin.Context, code int, msg string, data interface{}) {
	c.JSON(code, responseData{
		Code:    code,
		Message: msg,
		Data:    data,
	})
}

func FailureResponse(c *gin.Context, code int, msg string) {
	c.JSON(code, responseData{
		Code:    code,
		Message: msg,
	})
}
