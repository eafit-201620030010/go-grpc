# Go: Protocol Buffers & gRPC

High-performance microservices using gRPC in Go

## Prerequisites

- [Go dev](https://go.dev/)
- [Protocol Buffers](https://protobuf.dev/) & [How to install the protocol buffer compiler](https://grpc.io/docs/protoc-installation/)
- [Go plugins](https://grpc.io/docs/languages/go/quickstart/#prerequisites)

```bash
  $ go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
  $ go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
```

- [Docker](https://www.docker.com/)

## Tech Stack

- Go
- gRPC
- Protocol Baffers
- Docker
- Postgresql

## Deployment

To deploy this project run

### Git clone

```bash
  git clone https://github.com/eafit-201620030010/go-grpc.git
```

### Docker build

```bash
cd database/

docker build . -t jjchavarrg-grpc-db

docker run -p 54321:5432 jjchavarrg-grpc-db
```

### Run server-test

```bash
go run server-test/main.go
```

### Run server-student

```bash
go run server-student/main.go
```

### Run client or test in postman with gRPC

```bash
go run client/main.go
```

- Note: change the methods in client/main.go

### Regenerate gRPC code

- student proto

```bash
protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative studentpb/student.proto
```

- test proto

```bash
protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative testpb/test.proto
```

## Documentation

[Documentation](https://platzi.com/cursos/go-protobuffers-grpc/)

## License

[MIT](https://choosealicense.com/licenses/mit/)
