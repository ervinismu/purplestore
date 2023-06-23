package repository

import (
	"errors"
	"fmt"

	"github.com/ervinismu/purplestore/internal/app/model"
	"github.com/jmoiron/sqlx"
	log "github.com/sirupsen/logrus"
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

// delete category by id
func (cr *CategoryRepository) DeleteByID(id int) error {
	var (
		sqlStatement = `
			DELETE FROM categories
			WHERE id = $1
		`
	)

	result, err := cr.DB.Exec(sqlStatement, id)
	if err != nil {
		log.Error(fmt.Errorf("error CategoryRepository - DeleteByID : %w", err))
		return err
	}

	totalAffected, _ := result.RowsAffected()
	if totalAffected <= 0 {
		return errors.New("no record affected")
	}

	return nil
}

func (cr *CategoryRepository) Update(category model.Category) error {
	var (
		sqlStatement = `
			UPDATE categories
			SET updated_at = NOW(),
				name = $2,
				description = $3
			WHERE id = $1
		`
	)

	result, err := cr.DB.Exec(sqlStatement, category.ID, category.Name, category.Description)
	if err != nil {
		log.Error(fmt.Errorf("error CategoryRepository - UpdateByID : %w", err))
		return err
	}

	totalAffected, _ := result.RowsAffected()
	if totalAffected <= 0 {
		return errors.New("no record affected")
	}

	return nil
}
