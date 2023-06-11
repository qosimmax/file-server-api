package main

import (
	"context"
	"github.com/qosimmax/file-server-api/server"
	"log"
)

func main() {
	log.Println("Starting ...")

	ctx := context.Background()
	var s server.Server

	if err := s.Serve(ctx); err != nil {
		log.Fatal(err.Error())
	}

}
