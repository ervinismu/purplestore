package service

import "github.com/ervinismu/purplestore/internal/app/model"

type UserRepository interface {
	Create(user model.User) error
	Browse() ([]model.User, error)
	GetByEmailAndUsername(email string, username string) (model.User, error)
	GetByEmail(email string) (model.User, error)
	Update(user model.User) error
	DeleteByID(id string) error
	GetByID(userID int) (model.User, error)
}
