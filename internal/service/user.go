package service

import (
	"context"
	"errors"

	"github.com/adrianramadhan/synpasis-ecommerce-api/internal/dto"
	"github.com/adrianramadhan/synpasis-ecommerce-api/internal/model"
	"github.com/adrianramadhan/synpasis-ecommerce-api/internal/repository"
	"github.com/adrianramadhan/synpasis-ecommerce-api/internal/utils"
)

type UserService interface {
	Login(ctx context.Context, request dto.UserLoginRequest) (string, error)
	Register(ctx context.Context, request dto.UserRegisterRequest) error
}

type userService struct {
	userRepository repository.UserRepository
}

func NewUser(userRepository repository.UserRepository) UserService {
	return &userService{
		userRepository: userRepository,
	}
}

func (u *userService) Login(ctx context.Context, request dto.UserLoginRequest) (string, error) {
	user, err := u.userRepository.FindByEmail(ctx, request.Email)
	if err != nil {
		return "", err
	}

	if !utils.ComparePassword(user.Password, request.Password) {
		return "", errors.New("invalid password")
	}

	token, err := utils.GenerateToken(user.Id)

	if err != nil {
		return "", err
	}

	return token, nil
}

func (u *userService) Register(ctx context.Context, request dto.UserRegisterRequest) error {
	user, _ := u.userRepository.FindByEmail(ctx, request.Email)
	if user != nil {
		return errors.New("user already exists")
	}

	hashedPassword := utils.HashPassword(request.Password)
	newUser := model.User{
		Username: request.Username,
		Email:    request.Email,
		Password: hashedPassword,
	}

	err := u.userRepository.Create(ctx, newUser)
	if err != nil {
		return err
	}

	return nil
}
