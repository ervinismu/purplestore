package repository

import (
	"fmt"

	"github.com/ervinismu/purplestore/models"
	"github.com/ervinismu/purplestore/schema"
	"gorm.io/gorm"
)

type ProductRepository struct {
	DB *gorm.DB
}

func NewProductRepository(db *gorm.DB) ProductRepository {
	return ProductRepository{DB: db}
}

func (pr ProductRepository) GetProducts() ([]models.Product, error) {
	var products []models.Product
	err := pr.DB.Find(&products).Error
	if err != nil {
		return products, err
	}

	return products, nil
}

func (pr ProductRepository) ShowProduct(id string) (models.Product, error) {
	var product models.Product
	if err := pr.DB.Where("id = ?", id).First(&product).Error; err != nil {
		return product, err
	}

	return product, nil
}

func (pr ProductRepository) CreateProduct(data schema.CreateProductReq) (models.Product, error) {
	product := models.Product {
		Name: data.Name,
		Description: data.Description,
	}

	err := models.DB.Create(&product).Error
	if err != nil {
		return product, err
	}
	return product, nil
}

func (pr ProductRepository) UpdateProduct(id string, data schema.UpdateProductReq) (models.Product, error) {
	var product models.Product

	if err := pr.DB.Where("id = ?", id).First(&product).Error; err != nil {
		return product, err
	}

	product.Name = data.Name
	product.Description = data.Description

	fmt.Println(product)

	err := pr.DB.Save(&product).Error
	if err != nil {
		return product, err
	}
	return product, nil
}

func (pr ProductRepository) DeleteProduct(id string) (models.Product, error) {
	var product models.Product

	if err := pr.DB.Where("id = ?", id).First(&product).Error; err != nil {
		return product, err
	}

	err := pr.DB.Delete(&product).Error
	if err != nil {
		return product, err
	}
	return product, nil
}
