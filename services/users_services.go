package services

import (
	"github.com/nurzamanindra/bookstore_users-api/domain/users"
	"github.com/nurzamanindra/bookstore_users-api/utils/crypto_utils"
	"github.com/nurzamanindra/bookstore_users-api/utils/date_utils"
	"github.com/nurzamanindra/bookstore_users-api/utils/errors"
)

var (
	UserService userServiceInterface = &userService{}
)

type userService struct{}

type userServiceInterface interface {
	GetUser(int64) (*users.User, *errors.RestErr)
	CreateUser(users.User) (*users.User, *errors.RestErr)
	UpdateUser(users.User) (*users.User, *errors.RestErr)
	DeleteUser(int64) *errors.RestErr
	FindByStatus(string) (users.Users, *errors.RestErr)
}

func (s *userService) GetUser(userId int64) (*users.User, *errors.RestErr) {
	result := users.User{Id: userId}
	if err := result.Get(); err != nil {
		return nil, err
	}
	return &result, nil
}

func (s *userService) CreateUser(user users.User) (*users.User, *errors.RestErr) {
	if err := user.Validate(); err != nil {
		return nil, err
	}
	user.Status = users.StatusActive
	user.DateCreated = date_utils.GetNowDBFormat()
	user.Password = crypto_utils.GetMd5(user.Password)
	if err := user.Save(); err != nil {
		return nil, err
	}

	return &user, nil
}

func (s *userService) UpdateUser(user users.User) (*users.User, *errors.RestErr) {
	current, err := s.GetUser(user.Id)
	if err != nil {
		return nil, err
	}

	if user.FirstName != "" {
		current.FirstName = user.FirstName
	}
	if user.LastName != "" {
		current.LastName = user.LastName
	}
	if user.Email != "" {
		current.Email = user.Email
	}
	if err := current.Validate(); err != nil {
		return nil, err
	}

	if err := current.Update(); err != nil {
		return nil, err
	}
	return current, nil
}

func (s *userService) DeleteUser(userId int64) *errors.RestErr {
	current := &users.User{Id: userId}
	return current.Delete()
}

func (s *userService) FindByStatus(status string) (users.Users, *errors.RestErr) {
	dao := &users.User{}
	return dao.FindByStatus(status)
}
