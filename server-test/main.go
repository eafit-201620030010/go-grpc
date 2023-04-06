package main

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"jjchavarrg.com/go/grpc/database"
	"jjchavarrg.com/go/grpc/server"
	"jjchavarrg.com/go/grpc/testpb"
	"log"
	"net"
)

func main() {
	list, err := net.Listen("tcp", ":5070")

	if err != nil {
		log.Fatalf("Error listening: %s", err.Error())
	}

	repo, err := database.NewPostgresRepository("postgres://postgres:postgres@localhost:54321/postgres?sslmode=disable")

	server := server.NewTestServer(repo)

	if err != nil {
		log.Fatalf("Error creating repository: %s", err.Error())
	}

	s := grpc.NewServer()
	testpb.RegisterTestServiceServer(s, server)

	reflection.Register(s)

	if err := s.Serve(list); err != nil {
		log.Fatalf("Error serving: %s", err.Error())
	}
}
