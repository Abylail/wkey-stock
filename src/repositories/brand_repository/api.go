package brand_repository

import (
	"github.com/lowl11/lazy-entity/builders/select_builder"
	"wkey-stock/src/data/entities"
	"wkey-stock/src/data/models"
)

func (repo *Repository) UpdateIcon(id int, filePath string) error {
	return repo.UpdateByID(id, entities.Brand{
		Image: &filePath,
	})
}

func (repo *Repository) Create(model *models.BrandAdd) error {
	_, err := repo.Add(entities.Brand{
		Title: model.Title,
		Image: &model.Image,
	})
	if err != nil {
		return err
	}

	return nil
}

func (repo *Repository) Update(id int, model *models.BrandUpdate) error {
	return repo.UpdateByID(id, entities.Brand{
		Title: model.Title,
	})
}

func (repo *Repository) GetByTitle(title string) (*entities.Brand, error) {
	return repo.GetItem(func(builder *select_builder.Builder) {
		builder.Where(builder.Equal("title", title))
	})
}

func (repo *Repository) GetByQuery(searchQuery string) ([]entities.Brand, error) {
	searchQuery = "%" + searchQuery + "%"
	return repo.GetList(func(builder *select_builder.Builder) {
		builder.Where(builder.ILike("title", searchQuery))
	})
}
