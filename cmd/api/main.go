package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os/signal"
	"syscall"
	"time"

	"github.com/melnikdev/go-grafana/internal/config"
	"github.com/melnikdev/go-grafana/internal/database"
	"github.com/melnikdev/go-grafana/internal/server"
)

func main() {
	config := config.MustLoad()

	db := database.New(config.MongoDB)
	db.Health()

	defer func() {
		if err := db.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	server := server.NewServer(db, config.Server)

	done := make(chan bool, 1)

	go gracefulShutdown(server, done)

	err := server.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		panic(fmt.Sprintf("http server error: %s", err))
	}

	<-done
	log.Println("Graceful shutdown complete.")
}

func gracefulShutdown(apiServer *http.Server, done chan bool) {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	<-ctx.Done()

	log.Println("shutting down gracefully, press Ctrl+C again to force")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := apiServer.Shutdown(ctx); err != nil {
		log.Printf("Server forced to shutdown with error: %v", err)
	}

	log.Println("Server exiting")

	done <- true
}
