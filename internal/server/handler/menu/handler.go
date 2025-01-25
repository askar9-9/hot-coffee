package menu

import (
	"hot-coffee/internal/service"
	"hot-coffee/pkg/json"
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
	menuList, err := h.serv.GetAllMenuItems()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data, err := json.MarshalJson(menuList)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(data)
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
