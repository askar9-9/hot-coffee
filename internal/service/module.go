package service

import "hot-coffee/internal/entity"

type ServiceModule interface {
	InventoryItemService
	MenuService
}

type InventoryItemService interface {
	UpdateInventoryItemByID(item *entity.InventoryItem)
	GetInventoryItemByID() (*entity.InventoryItem, error)
	DeleteInventoryItemByID(item *entity.InventoryItem)
	GetAllInventoryItems() ([]entity.InventoryItem, error)
	AddInventoryItem(item *entity.InventoryItem)
}

type MenuService interface {
	GetAllMenuItems() ([]*entity.MenuItem, error)
	CreateMenuItem(item *entity.MenuItem) error

	GetMenuItemByID(id string) (*entity.MenuItem, error)
	ModifyMenuItem()
	RemoveMenuItem()

	CheckMenuItemValues(item *entity.MenuItem) error
}
