package server

import (
	"context"
	"fmt"
	"github.com/qosimmax/file-server-api/server/internal/handler"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
)

// Server holds the TCP server, router, config and all clients.
type Server struct {
	TcpListener net.Listener
}

func (srv *Server) Serve(ctx context.Context) error {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	var err error
	srv.TcpListener, err = net.Listen("tcp", ":"+port)
	if err != nil {
		return fmt.Errorf("error listening: %w", err)
	}

	log.Println("Started at port:" + port)

	go srv.acceptTcp(ctx)

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)

	<-stop

	log.Println("Shutdown signal received")

	return nil
}

func (srv *Server) acceptTcp(ctx context.Context) {
	for {
		conn, err := srv.TcpListener.Accept()
		if err != nil {
			log.Printf("Error accepting connection: %v\n", err)
		}

		go handler.HandleConnection(ctx, conn)
	}
}
