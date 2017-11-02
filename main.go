package main

import (
	"context"
	"log"
	"net"

	"github.com/roflomglol/estimator/maps"

	pb "github.com/roflomglol/estimator/estimation"
	"google.golang.org/grpc"
)

const (
	port = ":3001"
)

type server struct{}

func (s *server) TimeAndDistanceBetweenPoints(ctx context.Context, in *pb.PointsInfo) (*pb.DirectionsInfo, error) {
	start := maps.Point{Lat: in.Lat1, Long: in.Long1}
	finish := maps.Point{Lat: in.Lat2, Long: in.Long2}
	time, distance := maps.CalculateTimeAndDistance(&start, &finish)

	return &pb.DirectionsInfo{Time: int32(time), Distance: int32(distance)}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterEstimationServer(s, &server{})
	s.Serve(lis)
}
