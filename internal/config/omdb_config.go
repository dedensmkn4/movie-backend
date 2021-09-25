package config

import "os"

type OmdbConfig struct{
	BaseUrl string
	ApiKey string
}

func NewOmdbConfig() *OmdbConfig {
	omdbCfg := &OmdbConfig{}

	omdbCfg.BaseUrl = os.Getenv("OMDB_API_URL")
	omdbCfg.ApiKey = os.Getenv("OMDB_API_KEY")

	return omdbCfg
}