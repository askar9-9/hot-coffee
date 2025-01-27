package usecase

import (
	"fmt"
	"hot-coffee/internal/entity"
)

func (a *Application) AddInventoryItem(item *entity.InventoryItem) error {
	if err := validateInventoryItem(item); err != nil {
		return err
	}

	return a.repo.AddInventoryItem(item)
}

func (a *Application) GetAllInventoryItems() ([]*entity.InventoryItem, error) {
	return a.repo.GetAllInventoryItems()
}

func (a *Application) GetInventoryItemByID(id string) (*entity.InventoryItem, error) {
	return a.repo.GetInventoryItemByID(id)
}

func (a *Application) UpdateInventoryItemByID(id string, item *entity.InventoryItem) {

}

func (a *Application) DeleteInventoryItemByID(id string) {

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
