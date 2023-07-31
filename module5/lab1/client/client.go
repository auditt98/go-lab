package main

import (
	"context"
	"fmt"
	"log"

	pb "github.com/auditt98/go-lab/module5/lab1"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	c := pb.NewWeatherServiceClient(conn)
	req := &pb.WeatherRequest{
		Latitude:  7.857940,
		Longitude: 51.956055,
	}

	res, err := c.GetWeatherInfo(context.Background(), req)
	if err != nil {
		log.Fatalf("could not get weather information: %v", err)
	}

	fmt.Printf("Weather information: Temperature: %.2f\n", res.Temperature)
}
