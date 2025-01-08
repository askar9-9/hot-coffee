package server

import (
	"hot-coffee/internal/server/handler/inventory"
	"hot-coffee/internal/server/handler/menu"
	"hot-coffee/internal/service"
	"net/http"
)

func newRouter(serv service.ServiceModule) *http.ServeMux {
	router := http.NewServeMux()

	// Register routes
	menu.RegisterMenuRoutes(router, serv)
	inventory.RegisterInventoryRoutes(router, serv)

	return router
}
