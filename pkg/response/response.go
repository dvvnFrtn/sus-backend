package response

import (
	"net/http"
	"sus-backend/internal/dto"
	_error "sus-backend/pkg/err"

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
	code := mapErrorToStatusCode(err)
	c.JSON(code, dto.Response{
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

var errorToStatusCode = map[error]int{
	_error.ErrAlreadyLiked:   http.StatusBadRequest,
	_error.ErrNotLiked:       http.StatusBadRequest,
	_error.ErrNoOrganization: http.StatusBadRequest,
	_error.ErrNotFound:       http.StatusNotFound,
	_error.ErrConflict:       http.StatusConflict,
	_error.ErrUnauthorized:   http.StatusUnauthorized,
	_error.ErrForbidden:      http.StatusForbidden,
	_error.ErrInternal:       http.StatusInternalServerError,
}

func mapErrorToStatusCode(err error) int {
	if statusCode, exists := errorToStatusCode[err]; exists {
		return statusCode
	}
	return http.StatusInternalServerError
}
