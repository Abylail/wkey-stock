package sub_category_repository

import (
	"github.com/jmoiron/sqlx"
	"wkey-stock/src/data/entities"
)

const (
	statusActive   = "active"
	statusInactive = "inactive"
)

func (repo *Repository) setActive(parentID int, code, status string) error {
	ctx, cancel := repo.Ctx()
	defer cancel()

	entity := &entities.SubCategoryUpdateActive{
		Code:     code,
		ParentID: parentID,
		Status:   status,
	}

	query := repo.Script("sub_category", "update_active")

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
