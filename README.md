# movie-backend

## Made with
<p align="center">
  <img src="https://storage.googleapis.com/agungdp/static/logo/golang.png" width="80" alt="golang logo" />
  <img src="https://storage.googleapis.com/agungdp/static/logo/docker.png" width="80" hspace="10" alt="docker logo" />
  <img src="https://storage.googleapis.com/agungdp/static/logo/rest.png" width="80" hspace="10" alt="rest logo" />
  <img src="https://storage.googleapis.com/agungdp/static/logo/graphql.png" width="80" alt="graphql logo" />
  <img src="https://storage.googleapis.com/agungdp/static/logo/grpc.png" width="160" hspace="15" vspace="15" alt="grpc logo" />
  <img src="https://storage.googleapis.com/agungdp/static/logo/kafka.png" height="80" alt="kafka logo" />
</p>

This repository explain implementation of Go for building multiple microservices using a single codebase. Using [Standard Golang Project Layout](https://github.com/golang-standards/project-layout) and [Clean Architecture](https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html)

## How To Run Project

### Configuration
- Copy .env.sample to .env
- Create table from [here](./db)

### Install Package
``go mod tidy``

### Run GRPC Server
`` go run .\cmd\grpc-server\main.go``

### Run HTTP Server
`` go run .\cmd\http-server\main.go``

### Run Test
``make test``

## Api Doc
http://localhost:8092/swagger-ui