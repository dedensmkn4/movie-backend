package usecase

import (
	"github.com/dedensmkn4/movie-backend/internal/logging/domain"
	"github.com/dedensmkn4/movie-backend/internal/pb"
	mock "github.com/dedensmkn4/movie-backend/mock/repository"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"golang.org/x/net/context"
	"testing"
)

var (
	idMovieRequestSuccess = "tt4853102"
	pathFindByIdOmdbMovie = baseUrlOmdb+"/?apikey="+apiKey+"&i="+idMovieRequestSuccess
)

func TestOmdbService_FindById(t *testing.T) {

	testCases := []struct{
		testName	string
		context	context.Context
		getDetailRequest *pb.GetDetailRequest
		setupMockRepo setupMockRepo
		expectedErr	string
	}{
		{
			testName: "When search to omdb error store logging to db, return error",
			getDetailRequest: &pb.GetDetailRequest{
				Id: idMovieRequestSuccess,
			},
			setupMockRepo: func(mockRepo *mock.MockLoggingRepository) {
				mockRepo.EXPECT().
					InsertLoggingSearchOmdb(&domain.LoggingOmdbSearch{
						Path: pathFindByIdOmdbMovie,
						Response: successResponse,
					}).Return(int64(-1), errors.New(errorInsert))
			},
			expectedErr: errorInsert,
		},
		{
			testName: "When search to omdb success store logging to db, return success",
			getDetailRequest: &pb.GetDetailRequest{
				Id: idMovieRequestSuccess,
			},
			setupMockRepo: func(mockRepo *mock.MockLoggingRepository) {
				mockRepo.EXPECT().
					InsertLoggingSearchOmdb(&domain.LoggingOmdbSearch{
						Path: pathFindByIdOmdbMovie,
						Response: successResponse,
					}).Return(int64(1), nil)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			service, mockService := newOmdbServiceTest(t, tc.setupMockRepo)

			defer mockService.Finish()

			searchDetailResponse, err := service.FindById(tc.context, tc.getDetailRequest)

			if tc.expectedErr !="" {
				assert.EqualError(t, err, tc.expectedErr)
			}else {
				assert.NoError(t, err)
				assert.NotEmpty(t, searchDetailResponse)
			}
		})
	}


}
