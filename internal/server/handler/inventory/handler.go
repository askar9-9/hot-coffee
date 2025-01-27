package inventory

import (
	"hot-coffee/internal/entity"
	"hot-coffee/internal/server/handler"
	"hot-coffee/internal/service"
	"hot-coffee/pkg/json"
	"io"
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

	data, err := io.ReadAll(r.Body)
	defer r.Body.Close()

	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	newItem, err := json.UnmarshalJson[entity.InventoryItem](data)
	if err != nil {
		http.Error(w, "Invalid JSON format", http.StatusBadRequest)
		return
	}

	if err := h.serv.AddInventoryItem(&newItem); err != nil {
		http.Error(w, "Failed to add inventory item: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	_, _ = w.Write([]byte("Inventory item added successfully"))
}

func (h *InventoryHandler) GetInventoryItem(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	data, err := h.serv.GetInventoryItemByID(id)
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

func (h *InventoryHandler) UpdateInventoryItem(w http.ResponseWriter, r *http.Request) {
	if err := handler.Validation(w, r); err != nil {
		http.Error(w, err.Error(), http.StatusUnsupportedMediaType)
		return
	}

}

func (h *InventoryHandler) DeleteInventoryItem(w http.ResponseWriter, r *http.Request) {
	// Delete inventory item
}
