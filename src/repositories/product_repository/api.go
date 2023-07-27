package product_repository

import (
	"errors"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
	"strconv"
	"wkey-stock/src/data/entities"
	"wkey-stock/src/data/models"
)

func (repo *Repository) GetAdmin(from, pageSize int) ([]entities.AdminProductGet, error) {
	ctx, cancel := repo.Ctx()
	defer cancel()

	query := repo.Get("product", "get_admin")

	rows, err := repo.connection.QueryxContext(ctx, query, from, pageSize)
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

func (repo *Repository) GetAdminByQuery(from, pageSize int, searchQuery string) ([]entities.AdminProductGet, error) {
	ctx, cancel := repo.Ctx()
	defer cancel()

	searchQuery = "%" + searchQuery + "%"

	query := repo.Get("product", "get_admin_by_query")

	rows, err := repo.connection.QueryxContext(ctx, query, from, pageSize, searchQuery)
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

	query := repo.Get("product", "get_admin_by_id")

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

	query := repo.Get("product", "get_images")

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

	query := repo.Get("product", "count")

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

	query := repo.Get("product", "count_query")

	var count int
	if err := repo.connection.QueryRowxContext(ctx, query, searchQuery).Scan(&count); err != nil {
		return 0, err
	}

	return count, nil
}

func (repo *Repository) CountBySubCategory(subCategoryID int) (int, error) {
	ctx, cancel := repo.Ctx()
	defer cancel()

	query := repo.Get("product", "count_by_sub_category")

	var count int
	if err := repo.connection.QueryRowxContext(ctx, query, subCategoryID).Scan(&count); err != nil {
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

	query := repo.Get("product", "update")

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

func (repo *Repository) GetImagePositions(productID int, positions []int) ([]entities.ProductImageGet, error) {
	ctx, cancel := repo.Ctx()
	defer cancel()

	query := repo.Get("product", "get_image_positions")

	rows, err := repo.connection.QueryxContext(ctx, query, productID, pq.Array(positions))
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

func (repo *Repository) UpdateImages(id int, model *models.ProductUpload, pathList []string) error {
	if len(model.Images) != len(pathList) {
		return errors.New("images and theirs paths count does not match")
	}

	ctx, cancel := repo.Ctx()
	defer cancel()

	uploadEntities := make([]entities.ProductUpdateImage, 0, len(model.Images))

	for index, item := range model.Images {
		uploadEntities = append(uploadEntities, entities.ProductUpdateImage{
			ProductID: id,
			Path:      pathList[index],
			Position:  item.Position,
			Key:       strconv.Itoa(id) + "_" + strconv.Itoa(item.Position),
		})
	}

	query := repo.Get("product", "update_image")

	if err := repo.Transaction(repo.connection, func(tx *sqlx.Tx) error {
		for _, entity := range uploadEntities {
			if _, err := tx.NamedExecContext(ctx, query, entity); err != nil {
				return err
			}
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

	query := repo.Get("product", "bind")

	if err := repo.Transaction(repo.connection, func(tx *sqlx.Tx) error {
		for _, productID := range productIDs {
			if _, err := tx.ExecContext(ctx, query, productID, subCategoryID); err != nil {
				return err
			}
		}

		return nil
	}); err != nil {
		return err
	}

	return nil
}

func (repo *Repository) UnbindSubCategory(productID, subCategoryID int) error {
	ctx, cancel := repo.Ctx()
	defer cancel()

	query := repo.Get("product", "unbind")

	if err := repo.Transaction(repo.connection, func(tx *sqlx.Tx) error {
		if _, err := tx.ExecContext(ctx, query, productID, subCategoryID); err != nil {
			return err
		}

		return nil
	}); err != nil {
		return err
	}

	return nil
}

func (repo *Repository) GetSubCategoryPairs(productIDs []int) ([]entities.ProductCategoryPair, error) {
	ctx, cancel := repo.Ctx()
	defer cancel()

	query := repo.Get("product", "get_pairs")

	rows, err := repo.connection.QueryxContext(ctx, query, pq.Array(productIDs))
	if err != nil {
		return nil, err
	}
	defer repo.CloseRows(rows)

	list := make([]entities.ProductCategoryPair, 0)
	for rows.Next() {
		item := entities.ProductCategoryPair{}
		if err = rows.StructScan(&item); err != nil {
			return nil, err
		}
		list = append(list, item)
	}

	return list, nil
}

// GetClient список продуктов для клиента (без поискового текста)
func (repo *Repository) GetClient(from, pageSize int) ([]entities.ClientProductShort, error) {
	ctx, cancel := repo.Ctx()
	defer cancel()

	query := repo.Get("product", "get_client")

	rows, err := repo.connection.QueryxContext(ctx, query, from, pageSize)
	if err != nil {
		return nil, err
	}
	defer repo.CloseRows(rows)

	list := make([]entities.ClientProductShort, 0)
	for rows.Next() {
		item := entities.ClientProductShort{}
		if err = rows.StructScan(&item); err != nil {
			return nil, err
		}
		list = append(list, item)
	}

	return list, nil
}

// GetClientCount колличество товаров
func (repo *Repository) GetClientCount() (int, error) {
	ctx, cancel := repo.Ctx()
	defer cancel()

	query := repo.Get("product", "client_count")

	var count int
	if err := repo.connection.QueryRowxContext(ctx, query).Scan(&count); err != nil {
		return 0, err
	}

	return count, nil
}

// GetClientQuery список продуктов для клиента по поиску
func (repo *Repository) GetClientQuery(from, pageSize int, searchQuery string) ([]entities.ClientProductShort, error) {
	ctx, cancel := repo.Ctx()
	defer cancel()

	searchQuery = "%" + searchQuery + "%"

	query := repo.Get("product", "get_client_by_query")

	rows, err := repo.connection.QueryxContext(ctx, query, from, pageSize, searchQuery)
	if err != nil {
		return nil, err
	}
	defer repo.CloseRows(rows)

	list := make([]entities.ClientProductShort, 0)
	for rows.Next() {
		item := entities.ClientProductShort{}
		if err = rows.StructScan(&item); err != nil {
			return nil, err
		}
		list = append(list, item)
	}

	return list, nil
}

// GetClientCountQuery колличество товаров
func (repo *Repository) GetClientCountQuery(searchQuery string) (int, error) {
	ctx, cancel := repo.Ctx()
	defer cancel()

	searchQuery = "%" + searchQuery + "%"

	query := repo.Get("product", "get_client_count_by_query")

	var count int
	if err := repo.connection.QueryRowxContext(ctx, query, searchQuery).Scan(&count); err != nil {
		return 0, err
	}

	return count, nil
}
