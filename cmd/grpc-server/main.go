package main

import (
	"database/sql"
	"flag"
	"fmt"
	"github.com/dedensmkn4/movie-backend/internal/config"
	"github.com/dedensmkn4/movie-backend/internal/infra"
	"github.com/dedensmkn4/movie-backend/internal/logging/repository"
	"github.com/dedensmkn4/movie-backend/internal/omdb/usecase"
	"github.com/dedensmkn4/movie-backend/internal/pb"
	"github.com/dedensmkn4/movie-backend/internal/protocol/grpc/middleware"
	"github.com/dedensmkn4/movie-backend/pkg/logger"
	"github.com/joho/godotenv"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
	"strconv"
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

	logLevel, _ := strconv.Atoi(cfg.LogLevel)

	if err := logger.Init(logLevel, cfg.LogTimeFormat); err != nil {
		return errors.Wrap(err, "failed to initialize logger")
	}

	cfgDb := config.NewDatabaseConfig()
	db:= infra.NewDatabase(cfgDb)

	err = startGRPCServer(cfg, db.MySQL)
	if err != nil {
		return errors.Wrap(err, "Start GRPC Server")
	}

	return
}

func startGRPCServer(cfg *config.Config, db *sql.DB) (err error) {
	lis, err := net.Listen("tcp", cfg.GrpcServerAddr)
	if err != nil {
		log.Fatal(err)
	}

	var opts []grpc.ServerOption

	// add middleware
	opts = middleware.AddLogging(logger.Log, opts)

	gs := grpc.NewServer(opts...)

	omdbCfg := config.NewOmdbConfig()

	//register repo
	loggingRepo := repository.NewLoggingRepository(db)

	//register service
	omdbService := usecase.NewOmdbService(omdbCfg, loggingRepo)

	pb.RegisterOmdbServiceServer(gs, omdbService)
	if err := gs.Serve(lis); err != nil {
		return err
	}
	return
}