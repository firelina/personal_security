package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"golang.org/x/sync/errgroup"
	"log"
	"net/http"
	"os"
	"os/signal"
	httpGateway "personal_security/internal/gateways/http"
	"personal_security/internal/repository"
	"personal_security/internal/usecase"
	"syscall"
	"time"
)

func main() {
	server := http.Server{
		ReadHeaderTimeout: 10 * time.Second,
	}
	config, err := pgxpool.ParseConfig("postgres://postgres:postgres@127.0.0.1:5432/postgres?sslmode=disable")
	if err != nil {
		log.Fatalf("can't parse pgxpool config")
	}
	pool, err := pgxpool.NewWithConfig(context.Background(), config)
	if err != nil {
		log.Fatalf("can't create new pool")
	}
	defer pool.Close()
	userRepository := repository.NewUserRepository(pool)
	contactRepository := repository.NewContactRepository(pool)
	eventRepository := repository.NewEventRepository(pool)
	reminderRepository := repository.NewReminderRepository(pool)
	eventService := usecase.NewEventService(eventRepository)
	useCases := httpGateway.UseCases{
		User:     usecase.NewUserService(userRepository),
		Contact:  usecase.NewContactService(contactRepository),
		Event:    eventService,
		Reminder: usecase.NewReminderService(reminderRepository, eventService),
	}
	r := httpGateway.NewServer(useCases)
	server.Handler = r

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	eg, _ := errgroup.WithContext(context.Background())
	sigQuit := make(chan os.Signal, 1)
	signal.Notify(sigQuit, syscall.SIGINT, syscall.SIGTERM)
	eg.Go(func() error {
		s := <-sigQuit
		err := server.Shutdown(context.Background())
		if err != nil {
			log.Println(err.Error())
		}
		return fmt.Errorf("captured signal: %v", s)
	})

	go func() {
		if err := r.Run(ctx); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Printf("error during server shutdown: %v", err)
		}
	}()
	if err := eg.Wait(); err != nil {
		log.Printf("gracefully shutting down the server: %v", err) // gracefully shutting down the server: captured signal: interrupt
	}

}
