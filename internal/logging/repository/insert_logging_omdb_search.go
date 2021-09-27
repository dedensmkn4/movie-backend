package repository

import (
	"github.com/dedensmkn4/movie-backend/internal/logging/domain"
	"time"

	sq "github.com/Masterminds/squirrel"
)

func (l loggingRepositoryImpl) InsertLoggingSearchOmdb(ent *domain.LoggingOmdbSearch) (int64, error) {
	result, err := sq.
		Insert(domain.LoggingOmdbSearchTableName).
		Columns(
			domain.LoggingOmdbSearchTable.Path,
			domain.LoggingOmdbSearchTable.Response,
			domain.LoggingOmdbSearchTable.CreatedAt,
		).
		Values(
			ent.Path,
			ent.Response,
			time.Now(),
		).
		RunWith(l.db).
		Exec()

	if err != nil {
		return -1, err
	}
	return result.LastInsertId()
}