package repository

import (
	"database/sql"
	"github.com/dedensmkn4/movie-backend/internal/logging/domain"
)

type loggingRepositoryImpl struct {
	db	*sql.DB
}

type LoggingRepository interface {
	InsertLoggingSearchOmdb(ent *domain.LoggingOmdbSearch)(int64, error)
}

func NewLoggingRepository(db *sql.DB) LoggingRepository {
	return &loggingRepositoryImpl{
		db: db,
	}
}