package response

import (
	"sus-backend/internal/dto"

	"github.com/gin-gonic/gin"
)

func Success(c *gin.Context, httpCode int, msg string, data interface{}) {
	c.JSON(httpCode, dto.Response{
		Status:  true,
		Message: msg,
		Data:    data,
	})
}

func FailOrError(c *gin.Context, httpCode int, msg string, err error) {
	c.JSON(httpCode, dto.Response{
		Status:  false,
		Message: msg,
		Data:    nil,
	})
}

func ErrorEmptyField(c *gin.Context) {
	c.JSON(400, gin.H{
		"status":  "fail",
		"message": "Please fill the empty field",
	})
}
