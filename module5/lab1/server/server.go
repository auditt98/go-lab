package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net"
	"net/http"

	pb "github.com/auditt98/go-lab/module5/lab1"

	"google.golang.org/grpc"
)

type weatherServer struct {
	pb.UnimplementedWeatherServiceServer
}

func (s *weatherServer) GetWeatherInfo(ctx context.Context, req *pb.WeatherRequest) (*pb.WeatherResponse, error) {
	url := "https://api.openweathermap.org/data/2.5/weather" +
		"?lat=" + fmt.Sprintf("%.6f", req.Latitude) +
		"&lon=" + fmt.Sprintf("%.6f", req.Longitude) +
		"&units=metric&appid=81804a170a781fff7807ddc3eb8fb016"

	resp, err := http.Get(url)
	if err != nil {
		log.Printf("Failed to fetch weather data from OpenWeatherMap API: %v", err)
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Printf("OpenWeatherMap API returned non-OK status code: %v", resp.StatusCode)
		return nil, fmt.Errorf("failed to fetch weather data")
	}

	var weatherData struct {
		Main struct {
			Temp float64 `json:"temp"`
		} `json:"main"`
	}

	err = json.NewDecoder(resp.Body).Decode(&weatherData)
	if err != nil {
		log.Printf("Failed to decode weather data JSON: %v", err)
		return nil, err
	}

	temperature := weatherData.Main.Temp
	return &pb.WeatherResponse{
		Temperature: temperature,
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterWeatherServiceServer(s, &weatherServer{})
	log.Printf("Starting gRPC server on :50051")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
