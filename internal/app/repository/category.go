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
		data         []model.Category
		sqlStatement = "SELECT id, name, description FROM categories;"
	)
	rows, err := repo.DB.Queryx(sqlStatement)
	if err != nil {
		return data, err
	}
	for rows.Next() {
		var category model.Category
		rows.StructScan(&category)
		data = append(data, category)
	}

	return data, nil
}

func (repo *CategoryRepository) Create(category model.Category) error {
	sqlStatement := `
		INSERT INTO categories (name, descriptionss)
		VALUES ($1, $2)
	`
	_, err := repo.DB.Exec(sqlStatement, category.Name, category.Description)
	if err != nil {
		return err
	}

	return nil
}
