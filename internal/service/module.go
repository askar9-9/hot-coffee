package service

import "hot-coffee/internal/entity"

type ServiceModule interface {
	InventoryItemService
	MenuService
}

type InventoryItemService interface {
	UpdateInventoryItemByID(id string, item *entity.InventoryItem)
	GetInventoryItemByID(id string) (*entity.InventoryItem, error)
	DeleteInventoryItemByID(id string)
	GetAllInventoryItems() ([]*entity.InventoryItem, error)
	AddInventoryItem(item *entity.InventoryItem) error
}

type MenuService interface {
	GetAllMenuItems() ([]*entity.MenuItem, error)
	CreateMenuItem(item *entity.MenuItem) error
	GetMenuItemByID(id string) (*entity.MenuItem, error)
	ModifyMenuItem()
	RemoveMenuItem()
	CheckMenuItemValues(item *entity.MenuItem) error
}
