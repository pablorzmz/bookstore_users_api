package services

import (
	"github.com/pablorzmz/bookstore_users_api/src/domain/users"
	"github.com/pablorzmz/bookstore_users_api/src/utils/crypto_utils"
	"github.com/pablorzmz/bookstore_users_api/src/utils/date_utils"
	"github.com/pablorzmz/bookstore_users_api/src/utils/errors"
)

var (
	UserService userServiceInterface = &userService{}
)

type userService struct{}

type userServiceInterface interface {
	CreateUser(users.User) (*users.User, *errors.RestErr)
	GetUser(int64) (*users.User, *errors.RestErr)
	UpdateUser(bool, users.User) (*users.User, *errors.RestErr)
	DeleteUser(int64) *errors.RestErr
	FindByStatus(string) (users.Users, *errors.RestErr)
}

func (s *userService) CreateUser(user users.User) (*users.User, *errors.RestErr) {
	if err := user.Validate(); err != nil {
		return nil, err
	}

	user.DateCreated = date_utils.GetNowString()
	user.Status = users.StatusActive
	user.Password = crypto_utils.GetMD5(user.Password)

	if err := user.Save(); err != nil {
		return nil, err
	}

	return &user, nil
}

func (s *userService) GetUser(userId int64) (*users.User, *errors.RestErr) {
	result := users.User{Id: userId}
	if err := result.Get(); err != nil {
		return nil, err
	}
	return &result, nil
}
func (s *userService) UpdateUser(isPartial bool, user users.User) (*users.User, *errors.RestErr) {
	current, err := s.GetUser(user.Id)
	if err != nil {
		return nil, err
	}

	if err := user.Validate(); err != nil {
		return nil, err
	}

	if !isPartial {
		current.FirstName = user.FirstName
		current.Email = user.Email
		current.LastName = user.LastName
	} else {
		if user.FirstName != "" {
			current.FirstName = user.FirstName
		}
		if user.LastName != "" {
			current.LastName = user.LastName
		}
		if user.Email != "" {
			current.Email = user.Email
		}
	}

	err = current.Update()
	if err != nil {
		return nil, err
	}

	return current, nil
}

func (s *userService) DeleteUser(userId int64) *errors.RestErr {
	user := &users.User{Id: userId}
	return user.Delete()
}

func (s *userService) FindByStatus(status string) (users.Users, *errors.RestErr) {
	dao := &users.User{}
	return dao.FindByStatus(status)
}
