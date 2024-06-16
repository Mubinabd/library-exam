package main

import (
	"log"
	"net"

	pb "github.com/Mubinabd/library-exam/genproto"
	"github.com/Mubinabd/library-exam/service"
	"github.com/Mubinabd/library-exam/storage/postgres"
	"google.golang.org/grpc"
)

func main() {
	db, err := postgres.ConnectDB()
	if err != nil {
		panic(err)
	}
	liss, err := net.Listen("tcp", ":8085")
	if err != nil {
		panic(err)
	}
	
	s := grpc.NewServer()
	pb.RegisterAuthorServiceServer(s, service.NewAuthorStorage(db))
	pb.RegisterBookServiceServer(s, service.NewBookStorage(db))
	pb.RegisterBorrowerServiceServer(s, service.NewBorrowerStorage(db))
	pb.RegisterGenreServiceServer(s, service.NewGenreStorage(db))

	log.Printf("server listening at %v", liss.Addr())
	if err := s.Serve(liss); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
