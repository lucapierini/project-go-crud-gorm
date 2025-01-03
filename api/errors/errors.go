package errors

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

var (
	ErrBadRequest          = New(http.StatusBadRequest, "Bad request")
	ErrUnauthorized        = New(http.StatusUnauthorized, "Unauthorized")
	ErrNotFound            = New(http.StatusNotFound, "Resource not found")
	ErrInternalServerError = New(http.StatusInternalServerError, "Internal server error")
)


// APIError representa un error estructurado
type APIError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// New crea un nuevo APIError
func New(code int, message string) APIError {
	return APIError{Code: code, Message: message}
}

// RespondWithError envía un error al cliente
func RespondWithError(c *gin.Context, err APIError) {
	c.JSON(err.Code, gin.H{"error": err.Message})
}
