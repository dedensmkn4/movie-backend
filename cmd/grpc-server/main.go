package main

import (
	"flag"
	"fmt"
	"github.com/dedensmkn4/movie-backend/internal/config"
	"github.com/dedensmkn4/movie-backend/internal/omdb"
	"github.com/dedensmkn4/movie-backend/internal/pb"
	"github.com/joho/godotenv"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
)

func main(){

	if err := godotenv.Load(); err != nil {
		fmt.Println("error loading .env file")
	}

	flag.Parse()
	if err := run(); err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "%s/n", err)
		os.Exit(1)
	}
}

func run() (err error) {

	cfg := config.NewConfig()

	err = startGRPCServer(cfg)
	if err != nil {
		return errors.Wrap(err, "Start GRPC Server")
	}

	return
}

func startGRPCServer(cfg *config.Config) (err error) {
	lis, err := net.Listen("tcp", cfg.GrpcServerAddr)
	if err != nil {
		log.Fatal(err)
	}

	var opts []grpc.ServerOption

	gs := grpc.NewServer(opts...)

	omdbCfg := config.NewOmdbConfig()

	omdbService := omdb.NewOmdbService(omdbCfg)

	pb.RegisterOmdbServiceServer(gs, omdbService)
	if err := gs.Serve(lis); err != nil {
		return err
	}
	return
}