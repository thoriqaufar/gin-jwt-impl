package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/thoriqaufar/gin-jwt-impl/dto"
	"github.com/thoriqaufar/gin-jwt-impl/errorhandler"
	"github.com/thoriqaufar/gin-jwt-impl/helper"
	"github.com/thoriqaufar/gin-jwt-impl/service"
	"net/http"
)

type authHandler struct {
	service service.AuthService
}

func NewAuthHandler(s service.AuthService) *authHandler {
	return &authHandler{
		service: s,
	}
}

func (h *authHandler) Register(c *gin.Context) {
	var register dto.RegisterRequest

	if err := c.ShouldBindJSON(&register); err != nil {
		errorhandler.HandleError(c, &errorhandler.BadRequestError{Message: err.Error()})
		return
	}

	if err := h.service.Register(&register); err != nil {
		errorhandler.HandleError(c, err)
		return
	}

	response := helper.Response(dto.ResponseParams{
		StatusCode: http.StatusCreated,
		Message:    "Register successfully",
	})

	c.JSON(http.StatusCreated, response)
}

func (h *authHandler) Login(c *gin.Context) {
	var login dto.LoginRequest

	err := c.ShouldBindJSON(&login)
	if err != nil {
		errorhandler.HandleError(c, &errorhandler.BadRequestError{Message: err.Error()})
		return
	}

	result, err := h.service.Login(&login)
	if err != nil {
		errorhandler.HandleError(c, err)
		return
	}

	response := helper.Response(dto.ResponseParams{
		StatusCode: http.StatusOK,
		Message:    "Login successfully",
		Data:       result,
	})

	c.JSON(http.StatusOK, response)
}
