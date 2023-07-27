package promotion_repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/mehanizm/iuliia-go"
	"strings"
	"wkey-stock/src/data/entities"
	"wkey-stock/src/data/models"
)

// GetAll список все акций
func (repo *Repository) GetAll() ([]entities.AdminPromotion, error) {
	ctx, cancel := repo.Ctx()
	defer cancel()

	query := repo.Get("promotion", "get_all_admin")

	rows, err := repo.connection.QueryxContext(ctx, query)
	if err != nil {
		return nil, nil
	}

	list := make([]entities.AdminPromotion, 0)
	for rows.Next() {
		item := entities.AdminPromotion{}
		if err = rows.StructScan(&item); err != nil {
			return nil, err
		}
		list = append(list, item)
	}

	return list, nil
}

// GetById получить по id
func (repo *Repository) GetById(id int) (*entities.AdminPromotion, error) {
	ctx, cancel := repo.Ctx()
	defer cancel()

	query := repo.Get("promotion", "get_by_id")

	rows, err := repo.connection.QueryxContext(ctx, query, id)
	if err != nil {
		return nil, err
	}
	defer repo.CloseRows(rows)

	if rows.Next() {
		item := entities.AdminPromotion{}
		if err = rows.StructScan(&item); err != nil {
			return nil, err
		}
		return &item, nil
	}

	return nil, nil
}

// GetById получить по id
func (repo *Repository) GetByCode(code string) (*entities.AdminPromotion, error) {
	ctx, cancel := repo.Ctx()
	defer cancel()

	query := repo.Get("promotion", "get_by_code")

	rows, err := repo.connection.QueryxContext(ctx, query, code)
	if err != nil {
		return nil, err
	}
	defer repo.CloseRows(rows)

	if rows.Next() {
		item := entities.AdminPromotion{}
		if err = rows.StructScan(&item); err != nil {
			return nil, err
		}
		return &item, nil
	}

	return nil, nil
}

// New создать (возвращает код акции)
func (repo *Repository) Create(model *models.PromotionAdminCreate) (*string, error) {
	ctx, cancel := repo.Ctx()
	defer cancel()

	// генерируем код категории
	code := strings.TrimSpace(strings.ToLower(iuliia.Wikipedia.Translate(model.TitleRU)))
	code = strings.ReplaceAll(code, " ", "_")

	entity := &entities.AdminPromotionCreate{
		Code:          code,
		TitleRU:       model.TitleRU,
		TitleKZ:       model.TitleKZ,
		DescriptionRU: model.DescriptionRU,
		DescriptionKZ: model.DescriptionKZ,
	}

	query := repo.Get("promotion", "create")

	if err := repo.Transaction(repo.connection, func(tx *sqlx.Tx) error {
		if _, err := tx.NamedExecContext(ctx, query, entity); err != nil {
			return err
		}

		return nil
	}); err != nil {
		return nil, err
	}

	return &code, nil
}

// Update обновить акцию
func (repo *Repository) Update(model *models.PromotionAdminUpdate) error {
	ctx, cancel := repo.Ctx()
	defer cancel()

	entity := &entities.AdminPromotionUpdate{
		Code:          model.Code,
		TitleRU:       model.TitleRU,
		TitleKZ:       model.TitleKZ,
		DescriptionRU: model.DescriptionRU,
		DescriptionKZ: model.DescriptionKZ,
	}

	query := repo.Get("promotion", "update")

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

// UpdateImage загрузка картинки
func (repo *Repository) UpdateImage(code string, imagePath string, lang string) error {
	ctx, cancel := repo.Ctx()
	defer cancel()

	var imageField string = "image_ru"
	if lang == "kz" {
		imageField = "image_kz"
	}

	entity := &entities.AdminPromotionUpdateImage{
		Code:      code,
		ImagePath: imagePath,
	}

	query := fmt.Sprintf(repo.Get("promotion", "update_image"), imageField)

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

// Delete удалить акцию
func (repo *Repository) Delete(code *string) error {
	ctx, cancel := repo.Ctx()
	defer cancel()

	query := repo.Get("promotion", "delete")

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
