package inventory

import (
	"hot-coffee/internal/service"
	"net/http"
)

type InventoryHandler struct {
	serv service.ServiceModule
}

func NewInventoryHandler(serv service.ServiceModule) *InventoryHandler {
	return &InventoryHandler{
		serv: serv,
	}
}

func (h *InventoryHandler) ListInventoryItems(w http.ResponseWriter, r *http.Request) {
	// Get inventory items
}

func (h *InventoryHandler) AddInventoryItem(w http.ResponseWriter, r *http.Request) {
	// Add inventory item
}

func (h *InventoryHandler) GetInventoryItem(w http.ResponseWriter, r *http.Request) {
	// Get inventory item
}

func (h *InventoryHandler) UpdateInventoryItem(w http.ResponseWriter, r *http.Request) {
	// Update inventory item
}

func (h *InventoryHandler) DeleteInventoryItem(w http.ResponseWriter, r *http.Request) {
	// Delete inventory item
}
