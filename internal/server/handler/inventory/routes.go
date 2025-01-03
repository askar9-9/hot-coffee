package inventory

import (
	"hot-coffee/internal/service"
	"hot-coffee/pkg/logger"
	"net/http"
)

func RegisterInventoryRoutes(router *http.ServeMux, serv service.ServiceModule, logger *logger.CustomLogger) {
	inventoryHandler := NewInventoryHandler(serv, logger)

	router.HandleFunc("GET /inventory", inventoryHandler.ListInventoryItems)
	router.HandleFunc("POST /inventory", inventoryHandler.AddInventoryItem)
	router.HandleFunc("GET /inventory/{id}", inventoryHandler.GetInventoryItem)
	router.HandleFunc("PUT /inventory/{id}", inventoryHandler.UpdateInventoryItem)
	router.HandleFunc("DELETE /inventory/{id}", inventoryHandler.DeleteInventoryItem)
}
