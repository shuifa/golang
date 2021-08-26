package main

import (
	"context"
	"fmt"
	"log"

	pb "github/shuifa/golang/route"
	"google.golang.org/grpc"
)

func runFirst(client pb.RouteGuideClient) {
	feature, err := client.GetFeature(context.Background(), &pb.Point{
		Latitude:  31023500,
		Longitude: 121437403,
	})
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(feature)
}

func main() {
	conn, err := grpc.Dial("localhost:5000", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalln(err)
	}
	defer conn.Close()
	routeGuideClient := pb.NewRouteGuideClient(conn)
	runFirst(routeGuideClient)
}
