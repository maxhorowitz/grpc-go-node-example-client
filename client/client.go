package client

import (
	"time"

	pb "github.com/maxhorowitz/grpc-go-node-example-client"
)

func RequestData(stream pb.Messenger_SendClient) error {
	for i := 0; i < 5; i++ {
		if err := stream.Send(&pb.Request{Data: []byte("data")}); err != nil {
			return err
		}
		time.Sleep(time.Second * 2)
	}
	if err := stream.CloseSend(); err != nil {
		return err
	}
	return nil
}
