package postgres

import (
	"hot-coffee/pkg/logger"
)

type Postgres struct {
	log *logger.CustomLogger
}

func NewPostgres(log *logger.CustomLogger) *Postgres {
	return &Postgres{
		log: log,
	}
}
