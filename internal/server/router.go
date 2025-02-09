package server

import (
	"hot-coffee/internal/server/handler/menu"
	"hot-coffee/internal/service"
	"net/http"
)

func newRouter(serv service.Service) *http.ServeMux {
	router := http.NewServeMux()

	menu.RegisterRoutes(router, serv)
	return router
}
