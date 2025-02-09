package server

import (
	"hot-coffee/internal/server/middleware"
	"hot-coffee/internal/service"
	"hot-coffee/pkg/logger"
	"net/http"
)

type Server struct {
	addr   string
	router *http.ServeMux
	log    *logger.Logger
}

func NewServer(serv service.Service) *Server {
	router := newRouter(serv)

	addr := ":8080"
	return &Server{
		addr:   addr,
		router: router,
		log:    logger.GetLogger(),
	}
}

func (s *Server) Serve() error {
	wrappedRouter := middleware.Chain(s.router)
	return http.ListenAndServe(s.addr, wrappedRouter)
}
