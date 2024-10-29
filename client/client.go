package main

import (
	"context"
	"log"

	pb "github.com/amrshaban2005/grpc-protobuf/gen/go/hello/v1"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure(), grpc.WithBlock())

	if err != nil {
		log.Fatalf("could not connect to grpc server: %v", err)
	}

	defer conn.Close()

	client := pb.NewHelloServiceClient(conn)

	req := &pb.SayHelloRequest{
		Name: "Amr",
	}

	res, err := client.SayHello(context.Background(), req)

	if err != nil {
		log.Fatalf("Error calling SayHello: %v", err)
	}

	// Log the response
	log.Printf("Response from server: %s", res.GetMessage())

}
