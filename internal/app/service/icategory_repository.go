package service

import "github.com/ervinismu/purplestore/internal/app/model"

type CategoryRepository interface {
	GetList() ([]model.Category, error)
	Create(category model.Category) error
}
