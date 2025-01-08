package server

import (
	"hot-coffee/internal/server/middleware"
	"hot-coffee/internal/service"
	"hot-coffee/pkg/logger"
	"net/http"
	"strconv"
)

type Server struct {
	addr   string
	logger *logger.CustomLogger
	router *http.ServeMux
}

func NewServer(port int, serv service.ServiceModule, logger *logger.CustomLogger) *Server {
	router := newRouter(serv)

	addr := ":" + strconv.Itoa(port)
	return &Server{
		addr:   addr,
		router: router,
		logger: logger,
	}
}

func (s *Server) Serve() error {
	wrappedRouter := middleware.Chain(s.router,
		middleware.LoggingMiddleware(s.logger),
	)
	return http.ListenAndServe(s.addr, wrappedRouter)
}
