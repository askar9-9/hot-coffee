package server

import (
	"hot-coffee/internal/service"
	"hot-coffee/pkg/logger"
	"net/http"
	"strconv"
)

type Server struct {
	addr   string
	router *http.ServeMux
}

func NewServer(port int, serv service.ServiceModule, logger *logger.CustomLogger) *Server {
	router := newRouter(serv, logger)

	addr := ":" + strconv.Itoa(port)
	return &Server{
		addr:   addr,
		router: router,
	}
}

func (s *Server) Serve() error {
	return http.ListenAndServe(s.addr, s.router)
}
