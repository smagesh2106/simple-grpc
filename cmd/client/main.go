package main

import (
	"context"
	"log"
	"os"
	"time"

	pb "github.com/smagesh2106/simple-grpc/proto"
	"google.golang.org/grpc"
)

const (
	address     = "localhost:50051"
	defaultName = "world"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())

	if err != nil {
		log.Fatalf("did not connect :%v", err)
	}

	defer conn.Close()
	c := pb.NewGreeterClient(conn)

	name := defaultName
	if len(os.Args) > 1 {
		name = os.Args[1]
	}

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	//r, err := c.SayHello(ctx, &pb.HelloRequest{Name:name})
	r, err := c.SayHello(ctx, &pb.HelloRequest{Name: name})

	if err != nil {
		log.Fatalf("Could not greet :%v", err)
	} else {
		log.Printf("Response :%v", r.GetMessage())
	}
}
