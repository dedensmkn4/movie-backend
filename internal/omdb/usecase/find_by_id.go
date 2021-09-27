package usecase

import (
	"context"
	"encoding/json"
	"github.com/dedensmkn4/movie-backend/internal/pb"
)

func (o *OmdbService) FindById(_ context.Context, request *pb.GetDetailRequest) (*pb.SearchDetailResponse, error) {
	resp, err := o.callApi(searchId, request.GetId())

	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	r := new(pb.SearchDetailResponse)
	err = json.NewDecoder(resp.Body).Decode(r)

	if err != nil {
		return nil, err
	}

	return r, nil
}
