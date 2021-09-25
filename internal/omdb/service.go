package omdb

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/dedensmkn4/movie-backend/internal/config"
	"github.com/dedensmkn4/movie-backend/internal/pb"
	"github.com/pkg/errors"
	"net/http"
	"net/url"
	"strconv"
)

type OmdbService struct {
	pb.UnimplementedOmdbServiceServer
	omdbCfg *config.OmdbConfig
}

func NewOmdbService(omdbCfg *config.OmdbConfig) *OmdbService{
	return &OmdbService{omdbCfg: omdbCfg}
}


func (o *OmdbService) FindAll(ctx context.Context, request *pb.SearchRequest) (*pb.SearchPaginatedResponse, error) {
	pageParam := strconv.FormatInt(int64(request.GetPage()), 10)
	resp, err := o.callApi("search", request.GetSearchWord(), pageParam)

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

func (o *OmdbService) FindById(ctx context.Context, request *pb.GetDetailRequest) (*pb.SearchDetailResponse, error) {
	resp, err := o.callApi("id", request.GetId())

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

func (o *OmdbService)  callApi(apiCategory string, params...string) (resp *http.Response, err error) {
	var URL *url.URL
	URL, err = url.Parse(o.omdbCfg.BaseUrl)
	if err != nil {
		return nil, err
	}

	URL.Path += "/"
	parameters := url.Values{}
	parameters.Add("apikey", o.omdbCfg.ApiKey)

	switch apiCategory {
	case "search":
		parameters.Add("s", params[0])
		parameters.Add("page", params[1])
	case "id":
		parameters.Add("i", params[0])
	}

	URL.RawQuery = parameters.Encode()
	res, err := http.Get(URL.String())
	if err != nil {
		return nil, err
	}

	//if err=o.isValidResponse(res); err != nil{
	//	return nil, err
	//}
	fmt.Println(res)
	return res, nil
}

func (o *OmdbService) isValidResponse(request *http.Response)  error {

	r := new(OmdbOutboundDto)
	err := json.NewDecoder(request.Body).Decode(r)

	if err != nil {
		return err
	}
	fmt.Println(r)
	if r.Response == "False" {
		return errors.New(r.Error)
	}

	return nil
}