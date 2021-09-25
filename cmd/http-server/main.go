package main

import (
	"context"
	"fmt"
	"github.com/dedensmkn4/movie-backend/internal/config"
	"github.com/dedensmkn4/movie-backend/internal/pb"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/joho/godotenv"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"log"
	"net/http"
	"os"
)

func main()  {
	if err := run(); err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "%s/n", err)
		os.Exit(1)
	}
}

func run() (err error) {
	if err := godotenv.Load(); err != nil {
		fmt.Println("error loading .env file")
	}

	cfg := config.NewConfig()

	err = startHTTPServer(cfg)
	if err != nil {
		return errors.Wrap(err, "Start HTTP Server")
	}
	return
}

func startHTTPServer(cfg *config.Config) error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	conn, err := grpc.Dial(cfg.GrpcServerAddr, grpc.WithInsecure())
	if err != nil {
		return err
	}
	defer conn.Close()

	rmux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}
	err = pb.RegisterOmdbServiceHandlerFromEndpoint(ctx, rmux, cfg.GrpcServerAddr, opts)
	if err != nil {
		return err
	}

	// Swagger
	smux := http.NewServeMux()
	smux.Handle("/", rmux)
	smux.HandleFunc("/swagger.json", handleSwagger)
	fs := http.FileServer(http.Dir("www/swagger-ui"))
	smux.Handle("/swagger-ui/", http.StripPrefix("/swagger-ui", fs))

	log.Println("Running HTTP Server at " + cfg.HttpServerAddr)
	log.Println("Swagger is at: " + cfg.HttpServerAddr + "/swagger-ui/")
	err = http.ListenAndServe(cfg.HttpServerAddr, smux)
	if err != nil {
		return err
	}
	return nil
}

func handleSwagger(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "www/swagger.json")
}