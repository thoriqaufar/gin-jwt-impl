package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/thoriqaufar/gin-jwt-impl/dto"
	"github.com/thoriqaufar/gin-jwt-impl/errorhandler"
	"github.com/thoriqaufar/gin-jwt-impl/helper"
	"github.com/thoriqaufar/gin-jwt-impl/service"
	"net/http"
	"os"
	"path/filepath"
)

type postHandler struct {
	service service.PostService
}

func NewPostHandler(s service.PostService) *postHandler {
	return &postHandler{
		service: s,
	}
}

func (h *postHandler) Create(c *gin.Context) {
	var post dto.PostRequest

	if err := c.ShouldBind(&post); err != nil {
		errorhandler.HandleError(c, &errorhandler.BadRequestError{Message: err.Error()})
		return
	}

	if post.Picture != nil {
		if err := os.MkdirAll("/public/picture", 0755); err != nil {
			errorhandler.HandleError(c, &errorhandler.InternalServerError{Message: err.Error()})
			return
		}

		// Rename Picture
		ext := filepath.Ext(post.Picture.Filename)
		newFileName := uuid.New().String() + ext

		// Save image to directory
		dst := filepath.Join("public/picture", filepath.Base(newFileName))
		c.SaveUploadedFile(post.Picture, dst)

		post.Picture.Filename = fmt.Sprintf("%s/public/picture/%s", c.Request.Host, newFileName)
	}

	userID, _ := c.Get("userId")
	post.UserID = userID.(int)

	if err := h.service.Create(&post); err != nil {
		errorhandler.HandleError(c, err)
		return
	}

	response := helper.Response(dto.ResponseParams{
		StatusCode: http.StatusCreated,
		Message:    "Success post your tweet",
	})

	c.JSON(http.StatusCreated, response)
}
