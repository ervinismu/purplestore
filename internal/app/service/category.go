package service

import (
	"fmt"

	"github.com/ervinismu/purplestore/internal/app/model"
	"github.com/ervinismu/purplestore/internal/app/schema"
	log "github.com/sirupsen/logrus"
)

type CategoryRepository interface {
	GetList() ([]model.Category, error)
	Create(data model.Category) error
	GetByID(id int) (model.Category, error)
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
		errMsg := fmt.Errorf("category service - err get list %w", err)
		log.Error(errMsg)
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

func (svc *CategoryService) Create(req schema.CategoryCreateRequest) error {
	data := model.Category{
		Name:        req.Name,
		Description: req.Description,
	}
	err := svc.repo.Create(data)
	if err != nil {
		errMsg := fmt.Errorf("category service - err create %w", err)
		log.Error(errMsg)
		return err
	}

	return nil
}

func (svc *CategoryService) Detail(req schema.CategoryDetailRequest) (schema.CategoryDetailResponse, error) {
	var response schema.CategoryDetailResponse

	data, err := svc.repo.GetByID(req.ID)
	if err != nil {
		return response, err
	}

	response.ID = data.ID
	response.Name = data.Name
	response.Description = data.Description

	return response, nil
}
