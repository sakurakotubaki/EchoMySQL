package repository

import "myapi/internal/domain/model"

// ShoppingRepository defines the interface for shopping item storage
type ShoppingRepository interface {
	Create(item *model.ShoppingItem) error
	FindAll() ([]model.ShoppingItem, error)
	FindByID(id uint) (*model.ShoppingItem, error)
	Update(item *model.ShoppingItem) error
	Delete(id uint) error
}
