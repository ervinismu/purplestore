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

func (repo *CategoryRepository) Create(data model.Category) error {
	sqlStatement := `
		INSERT INTO categories (name, description)
		VALUES ($1, $2)
	`

	_, err := repo.DB.Exec(sqlStatement, data.Name, data.Description)
	if err != nil {
		return err
	}

	return nil
}

func (repo *CategoryRepository) GetByID(id int) (model.Category, error) {
	var (
		data         model.Category
		sqlStatement = `
			SELECT id, name, description
			FROM categories
			WHERE id = $1
			LIMIT 1
		`
	)

	err := repo.DB.QueryRowx(sqlStatement, id).StructScan(&data)
	if err != nil {
		return data, err
	}

	return data, nil
}
