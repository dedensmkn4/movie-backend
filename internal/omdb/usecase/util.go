package usecase

import (
	"bytes"
	"encoding/json"
	loggingDomain "github.com/dedensmkn4/movie-backend/internal/logging/domain"
	"github.com/dedensmkn4/movie-backend/internal/omdb/domain"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
)

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
	case searchPaginated:
		parameters.Add("s", params[0])
		parameters.Add("page", params[1])
	case searchId:
		parameters.Add("i", params[0])
	}

	URL.RawQuery = parameters.Encode()
	resp, err = http.Get(URL.String())
	if err != nil {
		return nil, err
	}

	_, err = o.loggingRepository.InsertLoggingSearchOmdb(o.toLoggingOmdbSearch(URL.String(), resp.Status))

	if err != nil {
		return nil, err
	}

	if err=o.isValidResponse(resp); err != nil{
		return nil, err
	}

	return
}

func (o *OmdbService) isValidResponse(request *http.Response)  error {

	res := new(domain.OmdbOutboundDto)

	b := bytes.NewBuffer(make([]byte, 0))
	reader := io.TeeReader(request.Body, b)

	err := json.NewDecoder(reader).Decode(res)

	if err != nil {
		return err
	}

	if res.Response == "False" {
		return status.Errorf(codes.InvalidArgument, res.Error)
	}

	request.Body = ioutil.NopCloser(b)

	return nil
}

func(o *OmdbService) toLoggingOmdbSearch(pathParam string, responseParam string) *loggingDomain.LoggingOmdbSearch{
	return &loggingDomain.LoggingOmdbSearch{
		Path: pathParam,
		Response: responseParam,
	}
}