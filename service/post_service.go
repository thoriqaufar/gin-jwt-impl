package service

import (
	"github.com/thoriqaufar/gin-jwt-impl/dto"
	"github.com/thoriqaufar/gin-jwt-impl/entity"
	"github.com/thoriqaufar/gin-jwt-impl/errorhandler"
	"github.com/thoriqaufar/gin-jwt-impl/repository"
)

type PostService interface {
	Create(request *dto.PostRequest) error
}

type postService struct {
	repository repository.PostRepository
}

func NewPostService(r repository.PostRepository) *postService {
	return &postService{
		repository: r,
	}
}

func (s *postService) Create(request *dto.PostRequest) error {
	post := entity.Post{
		UserID: request.UserID,
		Tweet:  request.Tweet,
	}

	if request.Picture != nil {
		post.PictureUrl = &request.Picture.Filename
	}

	if err := s.repository.Create(&post); err != nil {
		return &errorhandler.InternalServerError{Message: err.Error()}
	}

	return nil
}
