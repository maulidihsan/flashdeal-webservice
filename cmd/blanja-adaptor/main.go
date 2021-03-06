package main

import (
	"net"
	"log"
	"flag"
	"os"
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	murahid "github.com/maulidihsan/interop-commerce/cmd/blanja-adaptor/handler"
	"github.com/maulidihsan/interop-commerce/config"
)
func main() {
	environment := flag.String("e", "dev", "")
	flag.Usage = func() {
		fmt.Println("Usage: server -e {mode}")
		os.Exit(1)
	}
	flag.Parse()
	config.Init(*environment)
	c := config.GetConfig()

	listen, err := net.Listen("tcp", fmt.Sprintf(":%d", c.GetInt32("blanja.grpc.port")))
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	murahid.NewGrpcServer(grpcServer)
	reflection.Register(grpcServer)

	if err := grpcServer.Serve(listen); err != nil {
		log.Fatalf("Failed to serve: %s", err)
	}
}
