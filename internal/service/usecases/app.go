package usecases

import "hot-coffee/internal/storage"

type Application struct {
	repo storage.Repository
}

func NewApplication(repo storage.Repository) *Application {
	return &Application{
		repo: repo,
	}
}
