package category_repository

import (
	"github.com/jmoiron/sqlx"
	"wkey-stock/src/data/entities"
	"wkey-stock/src/data/models"
)

func (repo *Repository) GetAll() ([]entities.CategoryGet, error) {
	ctx, cancel := repo.Ctx()
	defer cancel()

	query := repo.Script("category", "get_all")

	rows, err := repo.connection.QueryxContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer repo.CloseRows(rows)

	list := make([]entities.CategoryGet, 0)
	for rows.Next() {
		item := entities.CategoryGet{}
		if err = rows.StructScan(&item); err != nil {
			return nil, err
		}
		list = append(list, item)
	}

	return list, nil
}

func (repo *Repository) GetByQuery(searchQuery string) ([]entities.CategoryGet, error) {
	ctx, cancel := repo.Ctx()
	defer cancel()

	searchQuery = "%" + searchQuery + "%"

	query := repo.Script("category", "get_by_query")

	rows, err := repo.connection.QueryxContext(ctx, query, searchQuery)
	if err != nil {
		return nil, err
	}
	defer repo.CloseRows(rows)

	list := make([]entities.CategoryGet, 0)
	for rows.Next() {
		item := entities.CategoryGet{}
		if err = rows.StructScan(&item); err != nil {
			return nil, err
		}
		list = append(list, item)
	}

	return list, nil
}

func (repo *Repository) GetByCode(code string) (*entities.CategoryGet, error) {
	ctx, cancel := repo.Ctx()
	defer cancel()

	query := repo.Script("category", "get_by_code")

	rows, err := repo.connection.QueryxContext(ctx, query, code)
	if err != nil {
		return nil, err
	}
	defer repo.CloseRows(rows)

	if rows.Next() {
		item := entities.CategoryGet{}
		if err = rows.StructScan(&item); err != nil {
			return nil, err
		}
		return &item, nil
	}

	return nil, nil
}

func (repo *Repository) Create(model *models.CategoryAdd, categoryCode string) error {
	ctx, cancel := repo.Ctx()
	defer cancel()

	entity := &entities.CategoryCreate{
		Code:    categoryCode,
		TitleRU: model.TitleRU,
		TitleKZ: model.TitleKZ,
	}

	query := repo.Script("category", "create")

	if err := repo.Transaction(repo.connection, func(tx *sqlx.Tx) error {
		if _, err := tx.NamedExecContext(ctx, query, entity); err != nil {
			return err
		}

		return nil
	}); err != nil {
		return err
	}

	return nil
}

func (repo *Repository) Update(code string, model *models.CategoryUpdate) error {
	ctx, cancel := repo.Ctx()
	defer cancel()

	entity := &entities.CategoryUpdate{
		Code:    code,
		TitleRU: model.TitleRU,
		TitleKZ: model.TitleKZ,
	}

	query := repo.Script("category", "update")

	if err := repo.Transaction(repo.connection, func(tx *sqlx.Tx) error {
		if _, err := tx.NamedExecContext(ctx, query, entity); err != nil {
			return err
		}

		return nil
	}); err != nil {
		return err
	}

	return nil
}

func (repo *Repository) UpdateImage(code string, imagePath string) error {
	ctx, cancel := repo.Ctx()
	defer cancel()

	entity := &entities.CategoryUpdateImage{
		Code:  code,
		Image: imagePath,
	}

	query := repo.Script("category", "update_image")

	if err := repo.Transaction(repo.connection, func(tx *sqlx.Tx) error {
		if _, err := tx.NamedExecContext(ctx, query, entity); err != nil {
			return err
		}

		return nil
	}); err != nil {
		return err
	}

	return nil
}

func (repo *Repository) Delete(code string) error {
	ctx, cancel := repo.Ctx()
	defer cancel()

	query := repo.Script("category", "delete")

	if err := repo.Transaction(repo.connection, func(tx *sqlx.Tx) error {
		if _, err := tx.ExecContext(ctx, query, code); err != nil {
			return err
		}

		return nil
	}); err != nil {
		return err
	}

	return nil
}
