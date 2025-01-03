package server

import (
	"hot-coffee/internal/server/handler/inventory"
	"hot-coffee/internal/server/handler/menu"
	"hot-coffee/internal/service"
	"hot-coffee/pkg/logger"
	"net/http"
)

func newRouter(serv service.ServiceModule, logger *logger.CustomLogger) *http.ServeMux {
	router := http.NewServeMux()

	// Register routes
	menu.RegisterMenuRoutes(router, serv, logger)
	inventory.RegisterInventoryRoutes(router, serv, logger)

	return router
}
