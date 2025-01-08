package menu

import (
	"hot-coffee/internal/service"
	"net/http"
)

func RegisterMenuRoutes(router *http.ServeMux, serv service.ServiceModule) {
	menuHandler := NewMenuHandler(serv)

	router.HandleFunc("GET /menu", menuHandler.ListMenuItems)
	router.HandleFunc("POST /menu", menuHandler.AddMenuItem)
	router.HandleFunc("GET /menu/{id}", menuHandler.GetMenuItem)
	router.HandleFunc("PUT /menu/{id}", menuHandler.UpdateMenuItem)
	router.HandleFunc("DELETE /menu/{id}", menuHandler.DeleteMenuItem)
}
