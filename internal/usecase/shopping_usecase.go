package usecase

import (
	"myapi/internal/domain/model"
	"myapi/internal/interface/repository"
)

// ShoppingUsecase defines the interface for shopping use cases
type ShoppingUsecase interface {
	CreateItem(name string) error
	GetAllItems() ([]model.ShoppingItem, error)
	GetItem(id uint) (*model.ShoppingItem, error)
	UpdateItem(id uint, name string) error
	DeleteItem(id uint) error
}

type shoppingUsecase struct {
	repo repository.ShoppingRepository
}

// NewShoppingUsecase creates a new shopping usecase
func NewShoppingUsecase(repo repository.ShoppingRepository) ShoppingUsecase {
	return &shoppingUsecase{repo: repo}
}

func (u *shoppingUsecase) CreateItem(name string) error {
	item := &model.ShoppingItem{Name: name}
	return u.repo.Create(item)
}

func (u *shoppingUsecase) GetAllItems() ([]model.ShoppingItem, error) {
	return u.repo.FindAll()
}

func (u *shoppingUsecase) GetItem(id uint) (*model.ShoppingItem, error) {
	return u.repo.FindByID(id)
}

func (u *shoppingUsecase) UpdateItem(id uint, name string) error {
	item, err := u.repo.FindByID(id)
	if err != nil {
		return err
	}
	item.Name = name
	return u.repo.Update(item)
}

func (u *shoppingUsecase) DeleteItem(id uint) error {
	return u.repo.Delete(id)
}
