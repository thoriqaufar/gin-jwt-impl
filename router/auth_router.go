package router

import (
	"github.com/gin-gonic/gin"
	"github.com/thoriqaufar/gin-jwt-impl/config"
	"github.com/thoriqaufar/gin-jwt-impl/handler"
	"github.com/thoriqaufar/gin-jwt-impl/repository"
	"github.com/thoriqaufar/gin-jwt-impl/service"
)

func AuthRouter(api *gin.RouterGroup) {
	authRepository := repository.NewAuthRepository(config.DB)
	authService := service.NewAuthService(authRepository)
	authHandler := handler.NewAuthHandler(authService)

	api.POST("/register", authHandler.Register)
	api.POST("/login", authHandler.Login)
}
