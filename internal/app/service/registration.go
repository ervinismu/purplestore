package service

import (
	"errors"

	"github.com/ervinismu/purplestore/internal/app/model"
	"github.com/ervinismu/purplestore/internal/app/schema"
	"github.com/ervinismu/purplestore/internal/pkg/reason"
	"golang.org/x/crypto/bcrypt"
)

type RegistrationService struct {
	userRepo UserRepository
}

func NewRegistrationService(userRepo UserRepository) *RegistrationService {
	return &RegistrationService{userRepo: userRepo}
}

func (svc *RegistrationService) Register(req *schema.RegisterRequest) error {
	existingUser, _ := svc.userRepo.GetByEmailAndUsername(req.Email, req.Username)
	if existingUser.ID > 0 {
		return errors.New(reason.RegisterFailed)
	}

	var insertData model.User
	password, _ := svc.hashPassword(req.Password)
	insertData.Username = req.Username
	insertData.Email = req.Email
	insertData.HashedPassword = password

	err := svc.userRepo.Create(insertData)
	if err != nil {
		return errors.New(reason.RegisterFailed)
	}

	return nil
}

func (svc *RegistrationService) hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(bytes), err
}
