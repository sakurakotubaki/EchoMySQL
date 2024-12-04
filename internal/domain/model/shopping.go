package model

import "github.com/jinzhu/gorm"

// ShoppingItem represents a shopping item in the system
type ShoppingItem struct {
	gorm.Model
	Name string `json:"name"`
}
