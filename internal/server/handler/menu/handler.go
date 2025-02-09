package menu

import (
	"hot-coffee/internal/server/handler/httperror"
	"hot-coffee/internal/service"
	"hot-coffee/pkg/logger"
	"net/http"
)

type MenuHanlder struct {
	serv service.MenuService
	log  *logger.Logger
}

func NewMenuHandler(serv service.MenuService) *MenuHanlder {
	return &MenuHanlder{
		serv: serv,
		log:  logger.GetLogger(),
	}
}

func (h *MenuHanlder) HandleRequest(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		h.GetMenu(w, r)
	case http.MethodPost:
		h.CreateMenu(w, r)
	default:
		httperror.MethodNotAllowed(w, r)
	}
}

func (h *MenuHanlder) HandleRequestByID(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		h.GetMenuByID(w, r)
	case http.MethodPut:
		h.UpdateMenuByID(w, r)
	case http.MethodDelete:
		h.DeleteMenuByID(w, r)
	default:
		httperror.MethodNotAllowed(w, r)
	}
}

// GET ["/menu"]
func (h *MenuHanlder) GetMenu(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("menu"))
}

// POST ["/menu"]
func (h *MenuHanlder) CreateMenu(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("create menu"))
}

// GET ["/menu/{id}"]
func (h *MenuHanlder) GetMenuByID(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("menu by id"))
}

// PUT ["/menu/{id}"]
func (h *MenuHanlder) UpdateMenuByID(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("update menu by id"))
}

// DELETE ["/menu/{id}"]
func (h *MenuHanlder) DeleteMenuByID(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("delete menu by id"))
}
