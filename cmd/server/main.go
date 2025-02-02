package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	httpGateway "personal_security/internal/gateways/http"
	"time"
)

func main() {
	server := http.Server{
		ReadHeaderTimeout: 10 * time.Second,
	}
	useCases := httpGateway.UseCases{}
	r := httpGateway.NewServer(useCases)
	server.Handler = r

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()
	if err := r.Run(ctx); err != nil {
		log.Printf("error during server shutdown: %v", err)
	}
}
