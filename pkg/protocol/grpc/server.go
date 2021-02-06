package grpc

import (
	"context"
	v1 "github.com/goblin2018/microservice_demo/pkg/api/v1"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
	"os/signal"
)

func RunServer(ctx context.Context, v1API v1.ToDoServiceServer, port string) error {
	listen, err := net.Listen("tcp", ":"+port)
	if err != nil {
		return err
	}

	server := grpc.NewServer()
	v1.RegisterToDoServiceServer(server, v1API)

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	go func() {
		for range c {
			log.Println("Shuttding down GRPC server ...")
			server.GracefulStop()
			<-ctx.Done()
		}
	}()

	log.Println("starting gRPC server ...")
	return server.Serve(listen)
}
