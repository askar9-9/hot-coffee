package main

import (
	"hot-coffee/internal/config"
	"hot-coffee/internal/server"
	"hot-coffee/internal/service/usecase"
	"hot-coffee/internal/storage/postgres"
	"hot-coffee/pkg/logger"
)

func main() {
	customLogger := logger.NewCustomLogger(
		config.InfoFilePath,
		config.ErrorFilePath,
		config.DebugFilePath,
	)

	repo := postgres.NewPostgres(customLogger)
	serv := usecase.NewApplication(repo, customLogger)
	
	serverHTTP := server.NewServer(config.Port, serv, customLogger)

	if err := serverHTTP.Serve(); err != nil {
		customLogger.FatalLogger.Fatalf("Failed to start server: %v", err)
	}
}
