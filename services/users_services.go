package services

import (
	"github.com/nurzamanindra/bookstore_users-api/domain/users"
	"github.com/nurzamanindra/bookstore_users-api/utils/errors"
)

func CreateUser(user users.User) (*users.User, *errors.RestErr) {
	if err := user.Validate(); err != nil {
		return nil, err
	}
	return &user, nil
}
