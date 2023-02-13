package brand_repository

import (
	"github.com/jmoiron/sqlx"
	"wkey-stock/src/data/entities"
	"wkey-stock/src/data/models"
)

func (repo *Repository) Create(model *models.BrandAdd) error {
	ctx, cancel := repo.Ctx()
	defer cancel()

	entity := &entities.BrandCreate{
		Title: model.Title,
		Image: model.Image,
	}

	query := repo.Script("brand", "create")

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

func (repo *Repository) Update(id int, model *models.BrandUpdate) error {
	ctx, cancel := repo.Ctx()
	defer cancel()

	entity := &entities.BrandUpdate{
		ID:    id,
		Title: model.Title,
	}

	query := repo.Script("brand", "update")

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

func (repo *Repository) UpdateIcon(id int, filePath string) error {
	ctx, cancel := repo.Ctx()
	defer cancel()

	entity := &entities.BrandUpload{
		ID:    id,
		Image: filePath,
	}

	query := repo.Script("brand", "update_icon")

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

func (repo *Repository) Delete(id int) error {
	ctx, cancel := repo.Ctx()
	defer cancel()

	query := repo.Script("brand", "delete")

	if err := repo.Transaction(repo.connection, func(tx *sqlx.Tx) error {
		if _, err := tx.ExecContext(ctx, query, id); err != nil {
			return err
		}

		return nil
	}); err != nil {
		return err
	}

	return nil
}

func (repo *Repository) GetAll() ([]entities.BrandGet, error) {
	ctx, cancel := repo.Ctx()
	defer cancel()

	query := repo.Script("brand", "get_all")

	rows, err := repo.connection.QueryxContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer repo.CloseRows(rows)

	list := make([]entities.BrandGet, 0)
	for rows.Next() {
		item := entities.BrandGet{}
		if err = rows.StructScan(&item); err != nil {
			return nil, err
		}
		list = append(list, item)
	}

	return list, nil
}

func (repo *Repository) GetByTitle(title string) (*entities.BrandGet, error) {
	ctx, cancel := repo.Ctx()
	defer cancel()

	query := repo.Script("brand", "get_by_title")

	rows, err := repo.connection.QueryxContext(ctx, query, title)
	if err != nil {
		return nil, err
	}
	defer repo.CloseRows(rows)

	if rows.Next() {
		item := entities.BrandGet{}
		if err = rows.StructScan(&item); err != nil {
			return nil, err
		}
		return &item, nil
	}

	return nil, nil
}

func (repo *Repository) GetByID(id int) (*entities.BrandGet, error) {
	ctx, cancel := repo.Ctx()
	defer cancel()

	query := repo.Script("brand", "get_by_id")

	rows, err := repo.connection.QueryxContext(ctx, query, id)
	if err != nil {
		return nil, err
	}
	defer repo.CloseRows(rows)

	if rows.Next() {
		item := entities.BrandGet{}
		if err = rows.StructScan(&item); err != nil {
			return nil, err
		}
		return &item, nil
	}

	return nil, nil
}

func (repo *Repository) GetByQuery(searchQuery string) ([]entities.BrandGet, error) {
	ctx, cancel := repo.Ctx()
	defer cancel()

	searchQuery = "%" + searchQuery + "%"

	query := repo.Script("brand", "get_by_query")

	rows, err := repo.connection.QueryxContext(ctx, query, searchQuery)
	if err != nil {
		return nil, err
	}
	defer repo.CloseRows(rows)

	list := make([]entities.BrandGet, 0)
	for rows.Next() {
		item := entities.BrandGet{}
		if err = rows.StructScan(&item); err != nil {
			return nil, err
		}
		list = append(list, item)
	}

	return list, nil
}
