package category_repository

import (
	"wkey-stock/src/data/entities"
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

func (repo *Repository) ByQuery(searchQuery string) ([]entities.CategoryGet, error) {
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
