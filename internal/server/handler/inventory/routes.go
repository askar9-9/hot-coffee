package inventory

import (
	"hot-coffee/internal/service"
	"net/http"
)

func RegisterInventoryRoutes(router *http.ServeMux, serv service.ServiceModule) {
	inventoryHandler := NewInventoryHandler(serv)

	router.HandleFunc("GET /inventory", inventoryHandler.ListInventoryItems)
	router.HandleFunc("POST /inventory", inventoryHandler.AddInventoryItem)
	router.HandleFunc("GET /inventory/{id}", inventoryHandler.GetInventoryItem)
	router.HandleFunc("PUT /inventory/{id}", inventoryHandler.UpdateInventoryItem)
	router.HandleFunc("DELETE /inventory/{id}", inventoryHandler.DeleteInventoryItem)
}
