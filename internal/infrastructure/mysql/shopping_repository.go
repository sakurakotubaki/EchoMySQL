package mysql

import (
	"myapi/internal/domain/model"

	"github.com/jinzhu/gorm"
)

type shoppingRepository struct {
	db *gorm.DB
}

// NewShoppingRepository creates a new shopping repository
func NewShoppingRepository(db *gorm.DB) *shoppingRepository {
	return &shoppingRepository{db: db}
}

func (r *shoppingRepository) Create(item *model.ShoppingItem) error {
	return r.db.Create(item).Error
}

func (r *shoppingRepository) FindAll() ([]model.ShoppingItem, error) {
	var items []model.ShoppingItem
	err := r.db.Find(&items).Error
	return items, err
}

func (r *shoppingRepository) FindByID(id uint) (*model.ShoppingItem, error) {
	var item model.ShoppingItem
	err := r.db.First(&item, id).Error
	if err != nil {
		return nil, err
	}
	return &item, nil
}

func (r *shoppingRepository) Update(item *model.ShoppingItem) error {
	return r.db.Save(item).Error
}

func (r *shoppingRepository) Delete(id uint) error {
	return r.db.Delete(&model.ShoppingItem{}, id).Error
}
