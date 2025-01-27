package usecase

import (
	"errors"
	"hot-coffee/internal/entity"
	"slices"
)

func (a *Application) GetAllMenuItems() ([]*entity.MenuItem, error) {
	return a.repo.GetAllMenuItems()
}

func (a *Application) CreateMenuItem(item *entity.MenuItem) error {
	return a.repo.CreateMenuItem(item)
}

func (a *Application) GetMenuItemByID(id string) (*entity.MenuItem, error) {
	return a.repo.GetMenuItemByID(id)
}

func (a *Application) ModifyMenuItem() {

}

func (a *Application) RemoveMenuItem() {

}

func (a *Application) CheckMenuItemValues(item *entity.MenuItem) error {
	if item == nil {
		return errors.New("menu item is required")
	}

	if item.Price <= 0 {
		return errors.New("price must be greater than 0")
	}

	if len(item.Name) < 3 {
		return errors.New("name must be at least 3 characters long")
	}

	if !slices.Contains([]string{"small", "medium", "large"}, item.Size) {
		return errors.New("invalid size")
	}

	// Check ingredients id

	return nil
}
