package main

import (
	"context"
	"fmt"
	"net"

	pb "grpc-pg/gen/worker"

	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedWorkerServiceServer
}

// this is the service that we are implementing that already has been defined in the gen/worker folder
func (s *server) SendHeartbeat(ctx context.Context, hb *pb.Heartbeat) (*pb.Ack, error) {
	fmt.Println(hb.WorkerId)
	return &pb.Ack{
		Message: "recieved",
	}, nil
}

func main() {
	//open a port at 50051 that accepts tcp connections(50051 just a convention)
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		panic(err)
	}
	//create a grpc server
	grpcServer := grpc.NewServer()

	//register a service
	pb.RegisterWorkerServiceServer(
		grpcServer,
		&server{},
	)
	//When somebody calls SendHeartbeat, execute methods on this server struct.
	//start server
	fmt.Println("gRPC server listening on :50051")

	if err := grpcServer.Serve(lis); err != nil {
		panic(err)
	}

}
