package service

import (
	"context"
	"github.com/Waelson/internal/model"
	"github.com/Waelson/internal/repository"
	"github.com/Waelson/internal/util/api"
	"github.com/Waelson/internal/util/encryptation"
	"github.com/Waelson/internal/util/token"
	"time"
)

type UserService interface {
	Save(ctx context.Context, user *model.User) (*model.User, api.Error)
	Login(ctx context.Context, login string, password string) (string, api.Error)
}

type userService struct {
	userRepository repository.UserRepository
}

func (s *userService) Save(ctx context.Context, user *model.User) (*model.User, api.Error) {

	userByLogin, err := s.userRepository.FindByLogin(ctx, user.Login)
	if err != nil {
		return nil, err
	}

	if userByLogin != nil {
		return nil, api.NewServiceError("login already exists")
	}

	userByEmail, err := s.userRepository.FindByEmail(ctx, user.Email)
	if err != nil {
		return nil, err
	}

	if userByEmail != nil {
		return nil, api.NewServiceError("email already exists")
	}

	hashedPass, errEnc := encryptation.Encrypt(user.Password)
	if errEnc != nil {
		return nil, api.NewInternalError(errEnc)
	}

	user.CreateAt = time.Now()
	user.Password = hashedPass

	_, err = s.userRepository.Save(ctx, user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *userService) Login(ctx context.Context, login string, password string) (string, api.Error) {
	userDb, err := s.userRepository.FindByLogin(ctx, login)
	if err != nil {
		return "", err
	}

	if userDb == nil || encryptation.Compare(userDb.Password, password) == false {
		return "", api.NewServiceError("login or password invalid")
	}

	result, errToken := token.Generate(login)
	if errToken != nil {
		return "", api.NewInternalError(errToken)
	}

	return result, nil
}

func NewUserService(userRepository repository.UserRepository) UserService {
	return &userService{
		userRepository: userRepository,
	}
}
