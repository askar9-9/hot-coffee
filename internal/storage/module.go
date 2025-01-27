package storage

import "hot-coffee/internal/entity"

type DataRepository interface {
	MenuRepository
	InventoryRepository
}

type MenuRepository interface {
	GetAllMenuItems() ([]*entity.MenuItem, error)
	GetMenuItemByID(id string) (*entity.MenuItem, error)
	CreateMenuItem(item *entity.MenuItem) error
	UpdateMenuItem(item *entity.MenuItem) error
	DeleteMenuItem(id string) error
}

type InventoryRepository interface {
	GetAllInventoryItems() ([]*entity.InventoryItem, error)
	GetInventoryItemByID(id string) (*entity.InventoryItem, error)
	AddInventoryItem(item *entity.InventoryItem) error
	UpdateInventoryItemByID(id string, item *entity.InventoryItem) error
	DeleteInventoryItemByID(id string) error
}
