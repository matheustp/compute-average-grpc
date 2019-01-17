package main

import (
	"context"
	"log"

	cavgpb "github.com/matheustp/compute-average-grpc/pb"
	"google.golang.org/grpc"
)

func main() {
	cc, err := grpc.Dial("0.0.0.0:50051", grpc.WithInsecure())
	if err != nil {
		log.Panicf("Error dialing server: %v", err)
	}
	c := cavgpb.NewComputeAverageServiceClient(cc)
	stream, err := c.ComputeAverage(context.Background())
	if err != nil {
		log.Panicf("Error calling function: %v", err)
	}

	numToCalc := []*cavgpb.ComputeAverageRequest{
		&cavgpb.ComputeAverageRequest{
			Num: 2,
		},
		&cavgpb.ComputeAverageRequest{
			Num: 2,
		},
		&cavgpb.ComputeAverageRequest{
			Num: 3,
		},
		&cavgpb.ComputeAverageRequest{
			Num: 3,
		},
	}

	for _, req := range numToCalc {
		log.Printf("Sending %v\n", req.GetNum())
		stream.Send(req)
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Panicf("Error when receiving response: %v", err)
	}
	log.Printf("Result: %v", res.GetResult())
}
