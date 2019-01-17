package main

import (
	"io"
	"log"
	"net"

	cavgpb "github.com/matheustp/compute-average-grpc/pb"
	"google.golang.org/grpc"
)

type server struct{}

func (*server) ComputeAverage(stream cavgpb.ComputeAverageService_ComputeAverageServer) error {
	var total, i int32
	for {
		res, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&cavgpb.ComputeAverageResponse{
				Result: float32(total) / float32(i),
			})
		}
		if err != nil {
			log.Panicf("Error when receiving request: %v", err)
		}
		total += res.GetNum()
		i++
	}
}

func main() {
	l, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Panicf("Error when starting to listen: %v", err)
	}
	s := grpc.NewServer()
	cavgpb.RegisterComputeAverageServiceServer(s, &server{})
	if err := s.Serve(l); err != nil {
		log.Panicf("Error when starting to serve: %v", err)
	}

}
