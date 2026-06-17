package main

import (
	"context"
	"fmt"
	"time"

	pb "grpc-pg/gen/worker"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() { //connect to the server
	conn, err := grpc.NewClient(
		"localhost:50051",
		grpc.WithTransportCredentials(
			insecure.NewCredentials(),
		),
	)
	defer conn.Close()
	if err != nil {
		panic(err)
	}
	// connect the client
	client := pb.NewWorkerServiceClient(conn)

	// build a req
	id := 1
	workerId := fmt.Sprintf("worker %d", id)
	hb := &pb.Heartbeat{
		WorkerId: workerId,
		CpuUsage: 42.5,
		RamUsage: 61.3,
		GpuUsage: 78.1,
	}

	// call the rpc
	for {
		resp, err := client.SendHeartbeat(
			context.Background(),
			hb,
		)
		id+=1
		time.Sleep(5 * time.Second)
		if err != nil {
			panic(err)
		}
		fmt.Println(resp.Message)
		fmt.Println(hb)
	}
}
