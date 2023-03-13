package middleware

import (
	"io"
	"strconv"
	service_error "todo_api/internal/service_errors"

	"github.com/faceair/jio"
	"github.com/gin-gonic/gin"
)

func ValidatePOSTBody(schema jio.Schema) gin.HandlerFunc {
	return func(c *gin.Context) {

		jsonData, err := io.ReadAll(c.Request.Body)
		if err != nil {
			c.Error(service_error.ServiceError("PARSE_JSON", "unable parse request json"))
			return
		}

		data, okData := jio.ValidateJSON(&jsonData, schema)

		if okData != nil {
			c.Error(service_error.ServiceError("VALIDATION_ERROR", okData.Error()))
			c.Abort()
			return
		}

		// Добавляем преобразованные данные в контекст запроса
		c.Set("requestBody", data)

		// Продолжаем выполнение следующего middleware или обработчика
		c.Next()
	}
}

func ValidateIdInQuery() gin.HandlerFunc {
	return func(c *gin.Context) {

		idStr := c.Param("id")
		id, okId := strconv.ParseUint(idStr, 10, 64)
		if okId != nil {
			c.Error(service_error.ServiceError("MUST_BE_NUMBER", "id must be a number"))
			c.Abort()
			return
		}

		// Добавляем преобразованные данные в контекст запроса
		c.Set("paramID", id)

		// Продолжаем выполнение следующего middleware или обработчика
		c.Next()
	}
}
