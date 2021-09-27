package usecase

import (
	"context"
	"github.com/dedensmkn4/movie-backend/internal/config"
	"github.com/dedensmkn4/movie-backend/internal/logging/domain"
	"github.com/dedensmkn4/movie-backend/internal/pb"
	mock "github.com/dedensmkn4/movie-backend/mock/repository"
	"github.com/golang/mock/gomock"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

var (
	apiKey = "faf7e5bb"
	baseUrlOmdb = "https://www.omdbapi.com/"
	searchSuccess = "batman"
	page int32= 2
	patchSearchOmdb = baseUrlOmdb + "/?apikey="+apiKey+"&page=2"+"&s="+searchSuccess
	successResponse="200 OK"
	errorInsert = "Insert error"
)

type setupMockRepo func(mockRepo *mock.MockLoggingRepository)

func newOmdbServiceTest(t *testing.T, fn setupMockRepo) (OmdbService, *gomock.Controller) {
	mockCtrl := gomock.NewController(t)

	loggingRepo := mock.NewMockLoggingRepository(mockCtrl)
	if fn!=nil {
		fn(loggingRepo)
	}

	omdbCfg := &config.OmdbConfig{}
	omdbCfg.ApiKey = apiKey
	omdbCfg.BaseUrl = baseUrlOmdb

	return OmdbService{
		UnimplementedOmdbServiceServer: pb.UnimplementedOmdbServiceServer{},
		omdbCfg :omdbCfg,
		loggingRepository:loggingRepo,
	}, mockCtrl
}


func TestOmdbService_FindAll(t *testing.T) {

	testCases := []struct{
		testName	string
		context context.Context
		searchRequest *pb.SearchRequest
		setupMockRepo setupMockRepo
		expectedErr	string
	}{
		{
			testName: "When search to omdb error store logging to db, return error",
			searchRequest: &pb.SearchRequest{
				SearchWord: searchSuccess,
				Page: page,
			},
			setupMockRepo: func(mockRepo *mock.MockLoggingRepository) {
				mockRepo.EXPECT().
					InsertLoggingSearchOmdb(&domain.LoggingOmdbSearch{
						Path: patchSearchOmdb,
						Response: successResponse,
				}).Return(int64(-1), errors.New(errorInsert))
			},
			expectedErr: errorInsert,
		},
		{
			testName: "When search to omdb success store logging to db, return success",
			searchRequest: &pb.SearchRequest{
				SearchWord: searchSuccess,
				Page: page,
			},
			setupMockRepo: func(mockRepo *mock.MockLoggingRepository) {
				mockRepo.EXPECT().
					InsertLoggingSearchOmdb(&domain.LoggingOmdbSearch{
						Path: patchSearchOmdb,
						Response: successResponse,
					}).Return(int64(1), nil)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			service, mockService := newOmdbServiceTest(t, tc.setupMockRepo)

			defer mockService.Finish()

			searchAllResponse, err := service.FindAll(tc.context, tc.searchRequest)

			if tc.expectedErr !="" {
				assert.EqualError(t, err, tc.expectedErr)
			}else {
				assert.NoError(t, err)
				assert.NotEmpty(t, searchAllResponse)
			}
		})
	}


}
