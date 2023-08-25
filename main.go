package main

import (
	"context"
	"math/rand"
	"os"

	"github.com/hashicorp/go-hclog"
	"google.golang.org/grpc"

	pb "github.com/maxhorowitz/grpc-go-node-example-api/pb/proto"
)

func main() {
	log := hclog.Default()
	// TODO (maxhorowitz): Do we want any dial options?
	// ANSWER: We do want dial options to set gRPC transport security. Need to learn about this.
	dialOpts := []grpc.DialOption{grpc.WithInsecure()}
	conn, err := grpc.Dial(":9092", dialOpts...)
	if err != nil {
		log.Error("Unable to dial connection", "error", err)
		os.Exit(1)
	}

	client := pb.NewMessengerClient(conn)
	callOpts := []grpc.CallOption{}
	stream, err := client.Connect(context.Background(), callOpts...)
	if err != nil {
		log.Error("Unable for client to call access", "error", err)
		os.Exit(1)
	}
	req := &pb.Req{
		Id:   int32(rand.Uint32()),
		Data: []byte("this is a test message"),
	}
	stream.Send(req)

}
