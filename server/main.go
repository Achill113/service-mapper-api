package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	pb "nordichill/servicemapper"
)

var (
	port = flag.Int("port", 50051, "The server port")
)

type server struct {
  pb.UnimplementedServiceMapperServer
}

func (s *server) CreateService(ctx context.Context, in *pb.CreateServiceRequest) (*pb.CreateServiceResponse, error) {
  log.Print("Creating service: %v", in.GetName())
  return &pb.CreateServiceResponse{Id: "123", Name: in.GetName()}, nil
}

func main() {
  flag.Parse()
  lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
  if err != nil {
    log.Fatalf("failed to listen: %v", err)
  }
  s := grpc.NewServer()
  pb.RegisterServiceMapperServer(s, &server{})
  log.Printf("server listening at %v", lis.Addr())
  if err := s.Serve(lis); err != nil {
    log.Fatalf("failed to serve: %v", err)
  }
}
