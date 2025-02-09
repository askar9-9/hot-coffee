package main

import (
	"hot-coffee/internal/server"
	"hot-coffee/internal/service/usecases"
	"hot-coffee/internal/storage/postgres"
	"hot-coffee/pkg/logger"
)

func main() {

	l := logger.GetLogger()
	
	repo := postgres.NewPostgres()
	serv := usecases.NewApplication(repo)
	s := server.NewServer(serv)

	if err := s.Serve(); err != nil {
		l.Fatalf("error serving: %v", err.Error())
	}
}
