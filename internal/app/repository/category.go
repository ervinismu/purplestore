package repository

import (
	"github.com/ervinismu/purplestore/internal/app/model"
	"github.com/jmoiron/sqlx"
)

type CategoryRepository struct {
	DB *sqlx.DB
}

func NewCategoryRepository(db *sqlx.DB) *CategoryRepository {
	return &CategoryRepository{DB: db}
}

func (repo *CategoryRepository) GetList() ([]model.Category, error) {
	var (
		categories   []model.Category
		sqlStatement = "SELECT id, name, description FROM categories"
	)

	rows, err := repo.DB.Queryx(sqlStatement)
	if err != nil {
		return categories, err
	}

	for rows.Next() {
		var category model.Category
		err := rows.StructScan(&category)
		if err != nil {
			return categories, err
		}
		categories = append(categories, category)
	}

	return categories, nil
}
