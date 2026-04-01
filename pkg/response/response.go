package response

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hogiabao7725/go-ticket-engine/pkg/apperror"
)

type SuccessResponse struct {
	Data any `json:"data"`
}

type ErrorResponse struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

func normalizeStatusCode(status int) int {
	if status < 100 || status > 599 {
		return http.StatusInternalServerError
	}
	return status
}

func OK(c *gin.Context, data any) {
	c.JSON(
		http.StatusOK,
		&SuccessResponse{Data: data},
	)
}

func Created(c *gin.Context, data any) {
	c.JSON(http.StatusCreated,
		&SuccessResponse{Data: data},
	)
}

func Error(c *gin.Context, err error) {
	var appErr *apperror.AppError

	if errors.As(err, &appErr) {
		c.JSON(normalizeStatusCode(appErr.StatusCode), ErrorResponse{
			Code:    appErr.Code,
			Message: appErr.Message,
		})
		return
	}

	c.JSON(http.StatusInternalServerError, ErrorResponse{
		Code:    apperror.ErrInternalServer.Code,
		Message: apperror.ErrInternalServer.Message,
	})
}
