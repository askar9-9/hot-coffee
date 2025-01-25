package usecase

import (
	"fmt"
	"hot-coffee/internal/entity"
)

func (a *Application) AddInventoryItem(item *entity.InventoryItem) {

}

func (a *Application) GetAllInventoryItems() ([]entity.InventoryItem, error) {
	// data , err := rep
	return []entity.InventoryItem{}, nil
}

func (a *Application) GetInventoryItemByID() (*entity.InventoryItem, error) {
	return &entity.InventoryItem{}, nil
}

func (a *Application) UpdateInventoryItemByID(item *entity.InventoryItem) {

}

func (a *Application) DeleteInventoryItemByID(item *entity.InventoryItem) {

}

func validateInventoryItem(item *entity.InventoryItem) error {
	if item.IngredientID == "" {
		return fmt.Errorf("ingredient ID is required")
	}
	if item.Name == "" {
		return fmt.Errorf("name is required")
	}
	if item.Quantity <= 0 {
		return fmt.Errorf("quantity must be greater than zero")
	}
	if item.Unit == "" {
		return fmt.Errorf("unit is required")
	}
	return nil
}
