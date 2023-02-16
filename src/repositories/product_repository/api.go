package product_repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
	"wkey-stock/src/data/entities"
	"wkey-stock/src/data/models"
)

func (repo *Repository) GetAdmin(from, to int) ([]entities.AdminProductGet, error) {
	ctx, cancel := repo.Ctx()
	defer cancel()

	query := repo.Script("product", "get_admin")

	rows, err := repo.connection.QueryxContext(ctx, query, from, to)
	if err != nil {
		return nil, err
	}
	defer repo.CloseRows(rows)

	list := make([]entities.AdminProductGet, 0)
	for rows.Next() {
		item := entities.AdminProductGet{}
		if err = rows.StructScan(&item); err != nil {
			return nil, err
		}
		list = append(list, item)
	}

	return list, nil
}

func (repo *Repository) GetAdminNoCategory(from, to int) ([]entities.AdminProductGet, error) {
	ctx, cancel := repo.Ctx()
	defer cancel()

	query := repo.Script("product", "get_admin_no_category")

	rows, err := repo.connection.QueryxContext(ctx, query, from, to)
	if err != nil {
		return nil, err
	}
	defer repo.CloseRows(rows)

	list := make([]entities.AdminProductGet, 0)
	for rows.Next() {
		item := entities.AdminProductGet{}
		if err = rows.StructScan(&item); err != nil {
			return nil, err
		}
		list = append(list, item)
	}

	return list, nil
}

func (repo *Repository) GetAdminByQuery(from, to int, searchQuery string) ([]entities.AdminProductGet, error) {
	ctx, cancel := repo.Ctx()
	defer cancel()

	searchQuery = "%" + searchQuery + "%"

	query := repo.Script("product", "get_admin_by_query")

	rows, err := repo.connection.QueryxContext(ctx, query, from, to, searchQuery)
	if err != nil {
		return nil, err
	}
	defer repo.CloseRows(rows)

	list := make([]entities.AdminProductGet, 0)
	for rows.Next() {
		item := entities.AdminProductGet{}
		if err = rows.StructScan(&item); err != nil {
			return nil, err
		}
		list = append(list, item)
	}

	return list, nil
}

func (repo *Repository) GetAdminNoCategoryByQuery(from, to int, searchQuery string) ([]entities.AdminProductGet, error) {
	ctx, cancel := repo.Ctx()
	defer cancel()

	searchQuery = "%" + searchQuery + "%"

	query := repo.Script("product", "get_admin_no_category_by_query")

	rows, err := repo.connection.QueryxContext(ctx, query, from, to, searchQuery)
	if err != nil {
		return nil, err
	}
	defer repo.CloseRows(rows)

	list := make([]entities.AdminProductGet, 0)
	for rows.Next() {
		item := entities.AdminProductGet{}
		if err = rows.StructScan(&item); err != nil {
			return nil, err
		}
		list = append(list, item)
	}

	return list, nil
}

func (repo *Repository) GetAdminByID(productID int) (*entities.AdminProductGet, error) {
	ctx, cancel := repo.Ctx()
	defer cancel()

	query := repo.Script("product", "get_admin_by_id")

	rows, err := repo.connection.QueryxContext(ctx, query, productID)
	if err != nil {
		return nil, err
	}
	defer repo.CloseRows(rows)

	if rows.Next() {
		item := entities.AdminProductGet{}
		if err = rows.StructScan(&item); err != nil {
			return nil, err
		}
		return &item, nil
	}

	return nil, nil
}

func (repo *Repository) GetImages(productIDs []int) ([]entities.ProductImageGet, error) {
	ctx, cancel := repo.Ctx()
	defer cancel()

	query := repo.Script("product", "get_images")

	rows, err := repo.connection.QueryxContext(ctx, query, pq.Array(productIDs))
	if err != nil {
		return nil, err
	}
	defer repo.CloseRows(rows)

	list := make([]entities.ProductImageGet, 0)
	for rows.Next() {
		item := entities.ProductImageGet{}
		if err = rows.StructScan(&item); err != nil {
			return nil, err
		}
		list = append(list, item)
	}

	return list, nil
}

func (repo *Repository) Count() (int, error) {
	ctx, cancel := repo.Ctx()
	defer cancel()

	query := repo.Script("product", "count")

	var count int
	if err := repo.connection.QueryRowxContext(ctx, query).Scan(&count); err != nil {
		return 0, err
	}

	return count, nil
}

func (repo *Repository) CountQuery(searchQuery string) (int, error) {
	ctx, cancel := repo.Ctx()
	defer cancel()

	searchQuery = "%" + searchQuery + "%"

	query := repo.Script("product", "count_query")

	var count int
	if err := repo.connection.QueryRowxContext(ctx, query, searchQuery).Scan(&count); err != nil {
		return 0, err
	}

	return count, nil
}

func (repo *Repository) CountNoCategory() (int, error) {
	ctx, cancel := repo.Ctx()
	defer cancel()

	query := repo.Script("product", "count_no_category")

	var count int
	if err := repo.connection.QueryRowxContext(ctx, query).Scan(&count); err != nil {
		return 0, err
	}

	return count, nil
}

func (repo *Repository) CountNoCategoryQuery(searchQuery string) (int, error) {
	ctx, cancel := repo.Ctx()
	defer cancel()

	searchQuery = "%" + searchQuery + "%"

	query := repo.Script("product", "count_no_category_query")

	var count int
	if err := repo.connection.QueryRowxContext(ctx, query, searchQuery).Scan(&count); err != nil {
		return 0, err
	}

	return count, nil
}

func (repo *Repository) Update(id int, model *models.ProductUpdate) error {
	ctx, cancel := repo.Ctx()
	defer cancel()

	entity := &entities.ProductUpdate{
		ID:            id,
		DescriptionRU: model.DescriptionRU,
		DescriptionKZ: model.DescriptionKZ,
	}

	query := repo.Script("product", "update")

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

func (repo *Repository) BindSubCategory(subCategoryID int, productIDs []int) error {
	ctx, cancel := repo.Ctx()
	defer cancel()

	query := repo.Script("product", "bind")

	if err := repo.Transaction(repo.connection, func(tx *sqlx.Tx) error {
		if _, err := tx.ExecContext(ctx, query, subCategoryID, productIDs); err != nil {
			return err
		}

		return nil
	}); err != nil {
		return err
	}

	return nil
}
