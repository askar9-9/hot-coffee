package menu

import (
	"hot-coffee/internal/service"
	"hot-coffee/pkg/logger"
	"net/http"
)

func RegisterMenuRoutes(router *http.ServeMux, serv service.ServiceModule, logger *logger.CustomLogger) {
	menuHandler := NewMenuHandler(serv, logger)

	router.HandleFunc("GET /menu", menuHandler.ListMenuItems)
	router.HandleFunc("POST /menu", menuHandler.AddMenuItem)
	router.HandleFunc("GET /menu/{id}", menuHandler.GetMenuItem)
	router.HandleFunc("PUT /menu/{id}", menuHandler.UpdateMenuItem)
	router.HandleFunc("DELETE /menu/{id}", menuHandler.DeleteMenuItem)
}
