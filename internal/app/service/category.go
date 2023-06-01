package service

import (
	"github.com/ervinismu/purplestore/internal/app/model"
	"github.com/ervinismu/purplestore/internal/app/schema"
)

type CategoryRepository interface {
	GetList() ([]model.Category, error)
}

type CategoryService struct {
	repo CategoryRepository
}

func NewCategorySerivce(repo CategoryRepository) *CategoryService {
	return &CategoryService{repo: repo}
}

func (svc *CategoryService) GetList() ([]schema.CategoryListResponse, error) {
	var response []schema.CategoryListResponse

	data, err := svc.repo.GetList()
	if err != nil {
		return response, err
	}

	for _, value := range data {
		var resp schema.CategoryListResponse
		resp.ID = value.ID
		resp.Name = value.Name
		resp.Description = value.Description
		response = append(response, resp)
	}

	return response, nil
}
