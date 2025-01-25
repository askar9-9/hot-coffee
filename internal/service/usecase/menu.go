package usecase

import "hot-coffee/internal/entity"

func (a *Application) GetAllMenuItems() ([]*entity.MenuItem, error) {
	return a.repo.GetAllMenuItems()
}

func (a *Application) CreateMenuItem() {

}

func (a *Application) GetMenuItemByID() {

}

func (a *Application) ModifyMenuItem() {

}

func (a *Application) RemoveMenuItem() {

}
