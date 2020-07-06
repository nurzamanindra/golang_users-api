package services

import (
	"github.com/nurzamanindra/golang_users-api/domain/users"
	"github.com/nurzamanindra/golang_users-api/utils/errors"
)

func CreateUser(user users.User) (*users.User, *errors.RestErr) {
	return &user, nil
}
