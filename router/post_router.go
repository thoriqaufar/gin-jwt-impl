package router

import (
	"github.com/gin-gonic/gin"
	"github.com/thoriqaufar/gin-jwt-impl/config"
	"github.com/thoriqaufar/gin-jwt-impl/handler"
	"github.com/thoriqaufar/gin-jwt-impl/middleware"
	"github.com/thoriqaufar/gin-jwt-impl/repository"
	"github.com/thoriqaufar/gin-jwt-impl/service"
)

func PostRouter(api *gin.RouterGroup) {
	postRepository := repository.NewPostRepository(config.DB)
	postService := service.NewPostService(postRepository)
	postHandler := handler.NewPostHandler(postService)

	r := api.Group("/tweets")

	r.Use(middleware.JWTMiddleware())

	r.POST("/", postHandler.Create)
}
