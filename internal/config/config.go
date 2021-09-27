package config

import "os"

type Config struct {
	GrpcServerAddr string
	HttpServerAddr string
	LogLevel string
	LogTimeFormat string
}

func NewConfig() *Config {
	cfg := &Config{}

	cfg.GrpcServerAddr = os.Getenv("APP_GRPC_SERVER_ADDR")
	cfg.HttpServerAddr = os.Getenv("APP_HTTP_SERVER_ADDR")
	cfg.LogLevel = os.Getenv("APP_LOG_LEVEL")
	cfg.LogTimeFormat = os.Getenv("APP_LOG_TIME_FORMAT")

	return cfg
}