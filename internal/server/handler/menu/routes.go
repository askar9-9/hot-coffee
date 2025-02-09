package menu

import (
	"hot-coffee/internal/service"
	"net/http"
)

func RegisterRoutes(router *http.ServeMux, serv service.Service) {
	h := NewMenuHandler(serv)

	router.HandleFunc("/menu", h.HandleRequest)
	router.HandleFunc("/menu/{id}", h.HandleRequestByID)
}
