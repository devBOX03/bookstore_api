package services

import (
	"github.com/devBOX03/bookstore_user_api/domain/users"
	"github.com/devBOX03/bookstore_user_api/utils/errors"
)

func CreateUser(user users.User) (*users.User, *errors.RestError) {
	if validateErr := user.Validate(); validateErr != nil {
		return nil, validateErr
	}

	daoErr := user.Save()
	if daoErr != nil {
		return nil, daoErr
	}
	return &user, nil
}

func GetUser(userId int64) (*users.User, *errors.RestError) {
	dao := &users.User{Id: userId}
	if err := dao.Get(); err != nil {
		return nil, err
	}
	return dao, nil
}

func UpdateUser(isPartial bool, user users.User) (*users.User, *errors.RestError) {

	currentUser := &users.User{Id: user.Id}
	if currentUserErr := currentUser.Get(); currentUserErr != nil {
		return nil, currentUserErr
	}

	if isPartial {
		if user.FirstName != "" {
			currentUser.FirstName = user.FirstName
		}
		if user.LastName != "" {
			currentUser.LastName = user.LastName
		}
		if user.Email != "" {
			currentUser.Email = user.Email
		}
	} else {
		currentUser.FirstName = user.FirstName
		currentUser.LastName = user.LastName
		currentUser.Email = user.Email
	}

	if updateErr := currentUser.Update(); updateErr != nil {
		return nil, updateErr
	}

	return currentUser, nil

}

func DeleteUser(userId int64) *errors.RestError {
	currentUser := &users.User{Id: userId}
	if currentUserErr := currentUser.Get(); currentUserErr != nil {
		return currentUserErr
	}
	return currentUser.Delete()
}

func SearchUser(status string) ([]users.User, *errors.RestError) {
	dao, daoErr := users.FindUserByStatus(status)
	if daoErr != nil {
		return nil, daoErr
	}
	return dao, nil

}
