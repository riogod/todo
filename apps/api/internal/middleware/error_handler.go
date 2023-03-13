package middleware

import (
	"net/http"
	service_error "todo_api/internal/service_errors"

	"github.com/gin-gonic/gin"
)

type ResponoseErrorDTO struct {
	Success bool                             `json:"success"`
	Error   service_error.ServiseErrorStruct `json:"error"`
}

func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {

		c.Next()

		if len(c.Errors) > 0 {
			// Get the last error
			err := c.Errors.Last().Err

			// Check if it's a ServiseError
			if serviceErr, ok := err.(*service_error.ServiseErrorStruct); ok {
				// Set the error response
				c.JSON(http.StatusOK, ResponoseErrorDTO{
					Success: false,
					Error:   *serviceErr,
				})

				// Abort the request
				c.Abort()
			}
		}

	}
}
