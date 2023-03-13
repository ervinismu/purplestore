package service

import (
	"github.com/ervinismu/purplestore/models"
	"github.com/ervinismu/purplestore/repository"
	"github.com/ervinismu/purplestore/schema"
)

type ProductService struct {
	repository repository.ProductRepository
}

func NewProductService(pr repository.ProductRepository) ProductService {
	return ProductService{repository: pr}
}

func (ps ProductService) ListAllProduct() ([]models.Product, error) {
	products, err := ps.repository.GetProducts()
	if err != nil {
		return nil, err
	}

	return products, nil
}

func (ps ProductService) ShowDetailProduct(id string) (models.Product, error) {
	product, err := ps.repository.ShowProduct(id)
	if err != nil {
		return product, err
	}

	return product, nil
}

func (ps ProductService) CreateProduct(data schema.CreateProductReq) (models.Product, error) {
	product, err := ps.repository.CreateProduct(data)
	if err != nil {
		return product, err
	}
	return product, nil
}

func (ps ProductService) DeleteProduct(id string) (models.Product, error) {
	product, err := ps.repository.DeleteProduct(id)
	if err != nil {
		return product, err
	}

	return product, nil
}

func (ps ProductService) UpdateProduct(id string, data schema.UpdateProductReq) (models.Product, error) {
	product, err := ps.repository.UpdateProduct(id, data)
	if err != nil {
		return product, err
	}

	return product, nil
}
