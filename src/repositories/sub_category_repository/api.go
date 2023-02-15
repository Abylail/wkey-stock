package sub_category_repository

import (
	"github.com/jmoiron/sqlx"
	"wkey-stock/src/data/entities"
	"wkey-stock/src/data/models"
)

func (repo *Repository) GetByParent(parentID int) ([]entities.SubCategoryGet, error) {
	ctx, cancel := repo.Ctx()
	defer cancel()

	query := repo.Script("sub_category", "get_by_parent")

	rows, err := repo.connection.QueryxContext(ctx, query, parentID)
	if err != nil {
		return nil, err
	}
	defer repo.CloseRows(rows)

	list := make([]entities.SubCategoryGet, 0)
	for rows.Next() {
		item := entities.SubCategoryGet{}
		if err = rows.StructScan(&item); err != nil {
			return nil, err
		}
		list = append(list, item)
	}

	return list, nil
}

func (repo *Repository) GetByQuery(parentID int, searchQuery string) ([]entities.SubCategoryGet, error) {
	ctx, cancel := repo.Ctx()
	defer cancel()

	searchQuery = "%" + searchQuery + "%"

	query := repo.Script("sub_category", "get_by_query")

	rows, err := repo.connection.QueryxContext(ctx, query, parentID, searchQuery)
	if err != nil {
		return nil, err
	}
	defer repo.CloseRows(rows)

	list := make([]entities.SubCategoryGet, 0)
	for rows.Next() {
		item := entities.SubCategoryGet{}
		if err = rows.StructScan(&item); err != nil {
			return nil, err
		}
		list = append(list, item)
	}

	return list, nil
}

func (repo *Repository) GetByCode(parentID int, code string) (*entities.SubCategoryGet, error) {
	ctx, cancel := repo.Ctx()
	defer cancel()

	query := repo.Script("sub_category", "get_by_code")

	rows, err := repo.connection.QueryxContext(ctx, query, parentID, code)
	if err != nil {
		return nil, err
	}
	defer repo.CloseRows(rows)

	if rows.Next() {
		item := entities.SubCategoryGet{}
		if err = rows.StructScan(&item); err != nil {
			return nil, err
		}
		return &item, nil
	}

	return nil, nil
}

func (repo *Repository) Create(parentID int, model *models.SubCategoryAdd, categoryCode string) error {
	ctx, cancel := repo.Ctx()
	defer cancel()

	entity := &entities.SubCategoryCreate{
		Code:     categoryCode,
		TitleRU:  model.TitleRU,
		TitleKZ:  model.TitleKZ,
		ParentID: parentID,
	}

	query := repo.Script("sub_category", "create")

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

func (repo *Repository) Update(parentID int, code string, model *models.SubCategoryUpdate) error {
	ctx, cancel := repo.Ctx()
	defer cancel()

	entity := &entities.SubCategoryUpdate{
		Code:     code,
		TitleRU:  model.TitleRU,
		TitleKZ:  model.TitleKZ,
		ParentID: parentID,
	}

	query := repo.Script("sub_category", "update")

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

	entity := &entities.SubCategoryUpdateImage{
		Code:  code,
		Image: imagePath,
	}

	query := repo.Script("sub_category", "update_image")

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

	query := repo.Script("sub_category", "delete")

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
