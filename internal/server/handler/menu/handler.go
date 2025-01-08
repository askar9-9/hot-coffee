package menu

import (
	"hot-coffee/internal/service"
	"net/http"
)

type MenuHandler struct {
	serv service.ServiceModule
}

func NewMenuHandler(serv service.ServiceModule) *MenuHandler {
	return &MenuHandler{
		serv: serv,
	}
}

func (h *MenuHandler) ListMenuItems(w http.ResponseWriter, r *http.Request) {
	// Get menu items
}

func (h *MenuHandler) AddMenuItem(w http.ResponseWriter, r *http.Request) {
	// Add menu item
}

func (h *MenuHandler) GetMenuItem(w http.ResponseWriter, r *http.Request) {
	// Get menu item
}

func (h *MenuHandler) UpdateMenuItem(w http.ResponseWriter, r *http.Request) {
	// Update menu item
}

func (h *MenuHandler) DeleteMenuItem(w http.ResponseWriter, r *http.Request) {
	// Delete menu item
}
