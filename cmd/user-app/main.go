package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"go-api-example/internal/app"
	"go-api-example/internal/port"
	"go-api-example/internal/storage"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

const (
	shutdownTimeoutSeconds = 60
)

func main() {
	listenAddr := flag.String("listenaddr", ":8080", "Listen address")
	flag.Parse()

	store := storage.NewMemoryStorage()

	application := app.NewApp(store)

	server, err := port.NewServer(*listenAddr, application)
	if err != nil {
		log.Fatalf("failed to start http server %s", err)
	}

	fmt.Printf("Listening on %s...\n", *listenAddr)
	go func() {
		if err = server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("Application start failed :: unable to start server %s", err)
		}
	}()

	<-signalChannel()

	shutdown(server.Shutdown)
}

func signalChannel() chan os.Signal {
	done := make(chan os.Signal, 1)
	signal.Notify(done, syscall.SIGINT, syscall.SIGTERM)
	return done
}

func shutdown(httpServer func(ctx context.Context) error) {
	log.Print("Application shutting down :: starts...")

	ctx, cancel := context.WithTimeout(context.Background(), shutdownTimeoutSeconds*time.Second)
	defer cancel()

	if err := httpServer(ctx); err != nil {
		log.Fatalf("Application shutting down failed :: unable to shut down http server %s", err)
	}

	log.Print("Application shutting down :: service shut down successfully")
}
