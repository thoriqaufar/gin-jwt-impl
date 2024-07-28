package service

import (
	"github.com/thoriqaufar/gin-jwt-impl/dto"
	"github.com/thoriqaufar/gin-jwt-impl/entity"
	"github.com/thoriqaufar/gin-jwt-impl/errorhandler"
	"github.com/thoriqaufar/gin-jwt-impl/helper"
	"github.com/thoriqaufar/gin-jwt-impl/repository"
)

type AuthService interface {
	Register(request *dto.RegisterRequest) error
	Login(request *dto.LoginRequest) (*dto.LoginResponse, error)
}

type authService struct {
	repository repository.AuthRepository
}

func NewAuthService(r repository.AuthRepository) *authService {
	return &authService{
		repository: r,
	}
}

func (s *authService) Register(request *dto.RegisterRequest) error {
	if emailExist := s.repository.EmailExists(request.Email); emailExist {
		return &errorhandler.BadRequestError{
			Message: "Email already exists",
		}
	}

	if request.Password != request.PasswordConfirmation {
		return &errorhandler.BadRequestError{
			Message: "Passwords do not match",
		}
	}

	passwordHash, err := helper.HashPassword(request.Password)
	if err != nil {
		return &errorhandler.InternalServerError{
			Message: err.Error(),
		}
	}

	user := entity.User{
		Name:     request.Name,
		Email:    request.Email,
		Password: passwordHash,
		Gender:   request.Gender,
	}

	if err := s.repository.Register(&user); err != nil {
		return &errorhandler.InternalServerError{
			Message: err.Error(),
		}
	}

	return nil
}

func (s *authService) Login(request *dto.LoginRequest) (*dto.LoginResponse, error) {
	var data dto.LoginResponse

	user, err := s.repository.GetUserByEmail(request.Email)
	if err != nil {
		return nil, &errorhandler.NotFoundError{Message: "Wrong email or password"}
	}

	if err := helper.VerifyPassword(user.Password, request.Password); err != nil {
		return nil, &errorhandler.NotFoundError{Message: "wrong email or password"}
	}

	token, err := helper.GenerateToken(user)
	if err != nil {
		return nil, &errorhandler.InternalServerError{Message: err.Error()}
	}

	data = dto.LoginResponse{
		ID:    user.ID,
		Name:  user.Name,
		Token: token,
	}

	return &data, nil
}
