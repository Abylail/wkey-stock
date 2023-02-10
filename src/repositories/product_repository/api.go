package product_repository

import "wkey-stock/src/data/entities"

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
