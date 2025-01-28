package menu

import (
	"errors"
	"hot-coffee/internal/entity"
	errors2 "hot-coffee/internal/errors"
	"hot-coffee/internal/server/handler"
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
	var data []byte

	if err := handler.Validation(w, r); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if _, err := r.Body.Read(data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	menuItem, err := json.UnmarshalJson[*entity.MenuItem](data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := h.serv.CheckMenuItemValues(menuItem); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := h.serv.CreateMenuItem(menuItem); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (h *MenuHandler) GetMenuItem(w http.ResponseWriter, r *http.Request) {
	// Get menu item
	var data []byte

	id := r.PathValue("id")

	menu, err := h.serv.GetMenuItemByID(id)
	if errors.Is(err, errors2.ErrMenuNotFound) {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data, err = json.MarshalJson(menu)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}

func (h *MenuHandler) UpdateMenuItem(w http.ResponseWriter, r *http.Request) {
	// Update menu item
}

func (h *MenuHandler) DeleteMenuItem(w http.ResponseWriter, r *http.Request) {
	// Delete menu item
}
