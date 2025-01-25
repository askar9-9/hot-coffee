package inventory

import (
	"hot-coffee/internal/server/handler"
	"hot-coffee/internal/service"
	"hot-coffee/pkg/json"
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

func (h *InventoryHandler) ListAllInventoryItems(w http.ResponseWriter, r *http.Request) {
	data, err := h.serv.GetAllInventoryItems()
	if err != nil {
		http.Error(w, "Error while get in data", http.StatusInternalServerError)
		return
	}

	responseData, err := json.MarshalJson(data)
	if err != nil {
		http.Error(w, "Error during data serialization", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(responseData)
}

func (h *InventoryHandler) AddInventoryItem(w http.ResponseWriter, r *http.Request) {
	if err := handler.Validation(w, r); err != nil {
		http.Error(w, err.Error(), http.StatusUnsupportedMediaType)
		return
	}

}

func (h *InventoryHandler) GetInventoryItem(w http.ResponseWriter, r *http.Request) {
	// Get inventory item
}

func (h *InventoryHandler) UpdateInventoryItem(w http.ResponseWriter, r *http.Request) {
	if err := handler.Validation(w, r); err != nil {
		http.Error(w, err.Error(), http.StatusUnsupportedMediaType)
		return
	}

}

func (h *InventoryHandler) DeleteInventoryItem(w http.ResponseWriter, r *http.Request) {
	// Delete inventory item
}
