package service

import "github.com/ervinismu/purplestore/internal/app/model"

type CategoryRepository interface {
	GetList(search model.CategorySearch) ([]model.Category, error)
	Create(data model.Category) error
	GetByID(id int) (model.Category, error)
	DeleteByID(id int) error
	Update(category model.Category) error
}
