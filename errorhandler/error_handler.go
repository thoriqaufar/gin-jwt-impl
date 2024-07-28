package errorhandler

import (
	"github.com/gin-gonic/gin"
	"github.com/thoriqaufar/gin-jwt-impl/dto"
	"github.com/thoriqaufar/gin-jwt-impl/helper"
	"net/http"
)

func HandleError(c *gin.Context, err error) {
	var statusCode int

	switch err.(type) {
	case *NotFoundError:
		statusCode = http.StatusNotFound
	case *BadRequestError:
		statusCode = http.StatusBadRequest
	case *InternalServerError:
		statusCode = http.StatusInternalServerError
	case *UnauthorizedError:
		statusCode = http.StatusUnauthorized
	}

	response := helper.Response(dto.ResponseParams{
		StatusCode: statusCode,
		Message:    err.Error(),
	})

	c.JSON(statusCode, response)
}
