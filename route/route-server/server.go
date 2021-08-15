package main

import (
	"context"
	"errors"
	"log"
	"net"

	pb "github/shuifa/route"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/proto"
)

type RouteGuideServer struct {
	Features []*pb.Feature
	pb.UnimplementedRouteGuideServer
}

func newServer() *RouteGuideServer {
	return &RouteGuideServer{
		Features: []*pb.Feature{
			{Name: "清华大学", Location: &pb.Point{
				Latitude:  310235000,
				Longitude: 121437403,
			}},
			{Name: "复旦大学", Location: &pb.Point{
				Latitude:  312978870,
				Longitude: 121503457,
			}},
			{Name: "深圳大学", Location: &pb.Point{
				Latitude:  311416130,
				Longitude: 121464904,
			}},
		},
	}
}

func (g *RouteGuideServer) GetFeature(ctx context.Context, point *pb.Point) (*pb.Feature, error) {
	for _, feature := range g.Features {
		if proto.Equal(point, feature.Location) {
			return feature, nil
		}
	}
	return nil, errors.New("gjgh")
}

func (g *RouteGuideServer) ListFeatures(*pb.Rectangle, pb.RouteGuide_ListFeaturesServer) error {
	return nil
}

func (g *RouteGuideServer) RecordRoute(pb.RouteGuide_RecordRouteServer) error {
	return nil
}

func (g *RouteGuideServer) Recommend(pb.RouteGuide_RecommendServer) error {
	return nil
}

func main() {
	listen, err := net.Listen("tcp", "localhost:5000")
	if err != nil {
		log.Fatalln(err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterRouteGuideServer(grpcServer, newServer())
	log.Fatalln(grpcServer.Serve(listen))
}
