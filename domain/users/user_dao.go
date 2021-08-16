package users

import (
	"fmt"

	"github.com/pablorzmz/bookstore_users_api/utils/errors"
)

var usersDB = make(map[int64]*User)

func (user *User) Save() *errors.RestErr {
	if current := usersDB[user.Id]; current != nil {
		if current.Email == user.Email {
			return errors.NewBadRequestError(fmt.Sprintf("Email %s already exists", user.Email))
		}
		return errors.NewBadRequestError("User already exists")
	}
	usersDB[user.Id] = user
	return nil
}

func (user *User) Get() *errors.RestErr {
	result := usersDB[user.Id]
	if result == nil {
		return errors.NotFoundError("User not found")
	}

	(*user).Id = result.Id
	user.FirstName = result.FirstName
	user.Email = result.Email
	user.LastName = result.LastName
	user.DateCreated = result.DateCreated

	return nil
}
