package order

import (
	"hot-coffee/internal/service"
	"net/http"
)

type OrderHandler struct {
	serv service.ServiceModule
}

func NewOrderHandler(serv service.ServiceModule) *OrderHandler {
	return &OrderHandler{
		serv: serv,
	}
}

func (h *OrderHandler) ListOrders(w http.ResponseWriter, r *http.Request) {
	// Get Orders
}

func (h *OrderHandler) AddNewOrder(w http.ResponseWriter, r *http.Request) {
	// Add New Order
}

func (h *OrderHandler) GetOrderById(w http.ResponseWriter, r *http.Request) {
	// Get Order ById
}

func (h *OrderHandler) UpdateOrderById(w http.ResponseWriter, r *http.Request) {
	// Update Order ById
}

func (h *OrderHandler) CloseOrder(w http.ResponseWriter, r *http.Request) {
	// Close Order ById
}

func (h *OrderHandler) DeleteOrderById(w http.ResponseWriter, r *http.Request) {
	// Delete Order ById
}
