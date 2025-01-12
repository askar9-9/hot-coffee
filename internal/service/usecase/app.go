package usecase

import (
	"hot-coffee/internal/storage"
)

type Application struct {
	repo storage.DataRepository
}

func NewApplication(repo storage.DataRepository) *Application {
	return &Application{
		repo: repo,
	}
}
