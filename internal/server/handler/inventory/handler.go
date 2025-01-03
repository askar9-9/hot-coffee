package inventory

import (
	"hot-coffee/internal/service"
	"hot-coffee/pkg/logger"
	"net/http"
)

type InventoryHandler struct {
	serv service.ServiceModule
	log  *logger.CustomLogger
}

func NewInventoryHandler(serv service.ServiceModule, log *logger.CustomLogger) *InventoryHandler {
	return &InventoryHandler{
		serv: serv,
		log:  log,
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
