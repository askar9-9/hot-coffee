package order

import (
	"hot-coffee/internal/service"
	"net/http"
)

func RegisterOrderRoutes(router *http.ServeMux, serv service.ServiceModule) {
	OrderHandler := NewOrderHandler(serv)

	router.HandleFunc("GET /orders", OrderHandler.ListOrders)
	router.HandleFunc("POST /order", OrderHandler.AddNewOrder)
	router.HandleFunc("GET /order/{id}", OrderHandler.GetOrderById)
	router.HandleFunc("PUT /order/{id}", OrderHandler.UpdateOrderById)
	router.HandleFunc("POST /order{id/close}", OrderHandler.CloseOrder)
	router.HandleFunc("DELETE /order/{id}", OrderHandler.DeleteOrderById)
}
