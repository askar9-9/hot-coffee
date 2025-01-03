package usecase

import (
	"hot-coffee/internal/storage"
	"hot-coffee/pkg/logger"
)

type Application struct {
	repo storage.DataRepository
	log  *logger.CustomLogger
}

func NewApplication(repo storage.DataRepository, log *logger.CustomLogger) *Application {
	return &Application{
		repo: repo,
		log:  log,
	}
}
