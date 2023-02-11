package brand_repository

import (
	"github.com/jmoiron/sqlx"
	"wkey-stock/src/data/entities"
	"wkey-stock/src/data/models"
)

func (repo *Repository) Create(model *models.BrandCreate) error {
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
