package service

import "hot-coffee/internal/entity"

type ServiceModule interface {
	InvenoryItemService
}

type InvenoryItemService interface {
	UpdateInventoryItemByID(item *entity.InventoryItem)
	GetInventoryItemByID() (*entity.InventoryItem, error)
	DeleteInventoryItemByID(item *entity.InventoryItem)
	GetAllInventoryItems() ([]entity.InventoryItem, error)
	AddInventoryItem(item *entity.InventoryItem)
}
