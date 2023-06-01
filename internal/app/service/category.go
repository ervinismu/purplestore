package service

import (
	"errors"

	"github.com/ervinismu/purplestore/internal/app/model"
	"github.com/ervinismu/purplestore/internal/app/schema"
	"github.com/ervinismu/purplestore/internal/pkg/reason"
)

type CategoryService struct {
	repo CategoryRepository
}

func NewCategoryService(repo CategoryRepository) *CategoryService {
	return &CategoryService{repo: repo}
}

func (svc *CategoryService) GetList() ([]schema.CategoryGetListResponse, error) {
	var response []schema.CategoryGetListResponse

	data, err := svc.repo.GetList()
	if err != nil {
		return response, err
	}

	for _, value := range data {
		var resp schema.CategoryGetListResponse
		resp.ID = value.ID
		resp.Name = value.Name
		resp.Description = value.Description
		response = append(response, resp)
	}

	return response, nil
}

func (svc *CategoryService) Create(req schema.CategoryCreateRequest) error {
	insertData := model.Category{}
	insertData.Name = req.Name
	insertData.Description = req.Description

	err := svc.repo.Create(insertData)
	if err != nil {
		return errors.New(reason.CategoryCreateFailed)
	}

	return nil
}
