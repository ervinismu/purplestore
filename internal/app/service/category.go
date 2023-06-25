package service

import (
	"errors"
	"fmt"

	"github.com/ervinismu/purplestore/internal/app/model"
	"github.com/ervinismu/purplestore/internal/app/schema"
	log "github.com/sirupsen/logrus"
)

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
		return response, errors.New("failed get list category")
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
		errMsg := fmt.Errorf("category service - err create : %w", err)
		log.Error(errMsg)
		return errors.New("failed create category")
	}

	return nil
}

func (svc *CategoryService) Detail(req schema.CategoryDetailRequest) (schema.CategoryDetailResponse, error) {
	var response schema.CategoryDetailResponse

	data, err := svc.repo.GetByID(req.ID)
	if err != nil {
		errMsg := fmt.Errorf("category service - err detail : %w", err)
		log.Error(errMsg)
		return response, errors.New("failed get detail category")
	}

	response.ID = data.ID
	response.Name = data.Name
	response.Description = data.Description

	return response, nil
}

func (cs *CategoryService) DeleteByID(req schema.CategoryDeleteRequest) error {

	_, err := cs.repo.GetByID(req.ID)
	if err != nil {
		return errors.New("failed delete category")
	}

	err = cs.repo.DeleteByID(req.ID)
	if err != nil {
		return errors.New("failed delete category")
	}

	return nil
}

func (cs *CategoryService) Update(req schema.CategoryUpdateRequest) error {

	var updateData model.Category

	oldData, err := cs.repo.GetByID(req.ID)
	if err != nil {
		return errors.New("failed update category")
	}

	updateData.ID = oldData.ID
	updateData.Name = req.Name
	updateData.Description = req.Description

	err = cs.repo.Update(updateData)
	if err != nil {
		return errors.New("failed update category")
	}

	return nil
}
