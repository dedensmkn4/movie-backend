package config

import "os"

type Config struct {
	GrpcServerAddr string
	HttpServerAddr string
}

func NewConfig() *Config {
	cfg := &Config{}

	cfg.GrpcServerAddr = os.Getenv("GRPC_SERVER_ADDR")
	cfg.HttpServerAddr = os.Getenv("HTTP_SERVER_ADDR")

	return cfg
}