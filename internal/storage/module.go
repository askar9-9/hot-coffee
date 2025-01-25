package storage

import "hot-coffee/internal/entity"

type DataRepository interface {
	MenuRepository
}

type MenuRepository interface {
	GetAllMenuItems() ([]*entity.MenuItem, error)
	GetMenuItemByID(id int) (*entity.MenuItem, error)
	CreateMenuItem(item *entity.MenuItem) error
	UpdateMenuItem(item *entity.MenuItem) error
	DeleteMenuItem(id int) error
}
