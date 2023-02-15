package category_repository

import (
	"github.com/jmoiron/sqlx"
	"wkey-stock/src/data/entities"
)

const (
	statusActive   = "active"
	statusInactive = "inactive"
)

func (repo *Repository) setActive(code, status string) error {
	ctx, cancel := repo.Ctx()
	defer cancel()

	entity := &entities.CategoryUpdateActive{
		Code:   code,
		Status: status,
	}

	query := repo.Script("category", "update_active")

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
