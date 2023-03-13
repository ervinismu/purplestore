package service

import (
	"github.com/ervinismu/purplestore/models"
	"github.com/ervinismu/purplestore/repository"
	"github.com/ervinismu/purplestore/schema"
)

type UserService struct {
	repository repository.UserRepository
}

func NewUserService(ur repository.UserRepository) UserService {
	return UserService{repository: ur}
}

func (ur UserService) ListAllUser() ([]models.User, error) {
	users, err := ur.repository.GetUsers()
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (ur UserService) ShowUser(id string) (models.User, error) {
	user, err := ur.repository.ShowUser(id)
	if err != nil {
		return user, err
	}

	return user, nil
}

func (ur UserService) CreateUser(data schema.CreateUserReq) (models.User, error) {
	user, err := ur.repository.CreateUser(data)
	if err != nil {
		return user, err
	}

	return user, nil
}

func (ur UserService) UpdateUser(id string, data schema.UpdateUserReq) (models.User, error) {
	user, err := ur.repository.UpdateUser(id, data)
	if err != nil {
		return user, err
	}

	return user, nil
}

func (ur UserService) DeleteUser(id string) (models.User, error) {
	user, err := ur.repository.DeleteUser(id)
	if err != nil {
		return user, err
	}

	return user, nil
}
