package responses

import (
	"github.com/gin-gonic/gin"
)

type responseData struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func Response(c *gin.Context, code int, msg string, data interface{}) {
	c.JSON(code, responseData{
		Code:    code,
		Message: msg,
		Data:    data,
	})
}
