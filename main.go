package main

import (
	"context"
	"flag"
	"log"
	"net"
	"net/http"
	"sync"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	pb "github.com/mhd999/weather-oracle/proto/go/services/weather"
	"github.com/mhd999/weather-oracle/weather"
	"google.golang.org/grpc"
)

var (
	grpcServerEndpoint = flag.String("grpc-server-endpoint", "127.0.0.1:50051", "gRPC server endpoint")
	grpcServerPort     = flag.String("grpc-server-port", ":50051", "gRPC server prot")
	jsonServerPort     = flag.String("json-server-port", ":8081", "JSON server port")
)

func init() {
	log.Printf("weather oracle GRPC is running on port: %s", *grpcServerPort)
	log.Printf("weather oracle JSON-API is running on port: %s", *jsonServerPort)
}

func runGRPCServer() error {
	lis, err := net.Listen("tcp", *grpcServerPort)
	if err != nil {
		return err
	}

	s := grpc.NewServer()

	pb.RegisterWeatherServer(s, &weather.Server{})

	return s.Serve(lis)
}

func runJSONServer() error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}
	err := pb.RegisterWeatherHandlerFromEndpoint(ctx, mux, *grpcServerEndpoint, opts)
	if err != nil {
		return err
	}

	return http.ListenAndServe(*jsonServerPort, mux)
}

func main() {
	wg := new(sync.WaitGroup)

	wg.Add(2)

	go func() {
		if err := runGRPCServer(); err != nil {
			log.Fatalf("failed to serve gRPC: %v", err)
		}

		wg.Done()
	}()

	go func() {
		if err := runJSONServer(); err != nil {
			log.Fatalf("failed to serve JSON-API: %v", err)
		}
		wg.Done()
	}()

	wg.Wait()

}
