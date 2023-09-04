package main

import (
	"context"
	"fmt"
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
	conn, err := grpc.Dial("localhost:8080", dialOpts...)
	if err != nil {
		log.Error("Unable to dial connection", "error", err)
		os.Exit(1)
	}

	registryClient := pb.NewRegistryClient(conn)
	callOpts := []grpc.CallOption{}

	firstName := &pb.FirstName{
		Name: "max",
	}
	lastName, err := registryClient.GetLast(context.Background(), firstName, callOpts...)
	log.Info(fmt.Sprint("received ", string(lastName.GetName())))
}
