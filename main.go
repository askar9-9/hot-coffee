package main

import (
	"hot-coffee/internal/server"
	"hot-coffee/internal/service/usecase"
	"hot-coffee/internal/storage/postgres"
	"hot-coffee/pkg/logger"
)

var (
	InfoFilePath  = "logs/info.log"
	ErrorFilePath = "logs/error.log"
	DebugFilePath = "logs/debug.log"
)

const (
	port = 8080
)

func main() {
	customLogger := logger.NewCustomLogger(
		InfoFilePath,
		ErrorFilePath,
		DebugFilePath,
	)

	repo := postgres.NewPostgres()
	serv := usecase.NewApplication(repo)

	serverHTTP := server.NewServer(port, serv, customLogger)

	if err := serverHTTP.Serve(); err != nil {
		customLogger.FatalLogger.Fatalf("Failed to start server: %v", err)
	}
}
