package main

import (
	"context"
	"log"
	"net"

	"github.com/roflomglol/estimator/maps"

	pb "github.com/roflomglol/estimator/estimator"
	"google.golang.org/grpc"
)

const (
	port = ":3001"
)

type server struct{}

func (s *server) TimeAndDistanceBetweenPoints(ctx context.Context, in *pb.PointsInfo) (*pb.DirectionsInfo, error) {
	origin := maps.Point{Lat: in.Origin.Lat, Long: in.Origin.Long}
	destination := maps.Point{Lat: in.Destination.Lat, Long: in.Destination.Long}
	time, distance := maps.CalculateTimeAndDistance(&origin, &destination)

	return &pb.DirectionsInfo{Time: time, Distance: distance}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterEstimateServer(s, &server{})
	s.Serve(lis)
}
