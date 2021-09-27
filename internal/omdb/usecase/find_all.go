package usecase

import (
	"context"
	"encoding/json"
	"github.com/dedensmkn4/movie-backend/internal/pb"
	"strconv"
)

func (o *OmdbService) FindAll(_ context.Context, request *pb.SearchRequest) (*pb.SearchPaginatedResponse, error) {

	pageParam := strconv.FormatInt(int64(request.GetPage()), 10)
	resp, err := o.callApi(searchPaginated, request.GetSearchWord(), pageParam)

	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	r := new(pb.SearchPaginatedResponse)
	err = json.NewDecoder(resp.Body).Decode(r)

	if err != nil {
		return nil, err
	}

	return r, nil
}

