package main

import (
	"context"
	"log"
	"net"

	pb "github.com/amrshaban2005/grpc-protobuf/gen/go/hello/v1"
	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedHelloServiceServer
}

func (s *server) SayHello(ctx context.Context, req *pb.SayHelloRequest) (*pb.SayHelloResponse, error) {
	log.Printf("recieved request from name: %s", req.GetName())
	return &pb.SayHelloResponse{
		Message: "Hello, " + req.GetName() + "!",
	}, nil

}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to list to port 50051: %v", err)

	}

	log.Println("grpc server listiening on port: 50051")

	grpcServer := grpc.NewServer()

	pb.RegisterHelloServiceServer(grpcServer, &server{})

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve gRPC server: %v", err)
	}

}
