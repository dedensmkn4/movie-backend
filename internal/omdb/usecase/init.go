package usecase

import (
	"github.com/dedensmkn4/movie-backend/internal/config"
	"github.com/dedensmkn4/movie-backend/internal/logging/repository"
	"github.com/dedensmkn4/movie-backend/internal/pb"
)

const (
	searchPaginated = "OMDB_SEARCH_PAGINATED"
	searchId        = "OMDB_SEARCH_ID"
)

type OmdbService struct {
	pb.UnimplementedOmdbServiceServer
	omdbCfg *config.OmdbConfig
	loggingRepository repository.LoggingRepository
}

func NewOmdbService(omdbCfg *config.OmdbConfig, loggingRepository repository.LoggingRepository) *OmdbService {
	return &OmdbService{
		omdbCfg: omdbCfg,
		loggingRepository: loggingRepository,
	}
}
