package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"

	api "github.com/voitenkov-courses/architecture-pro-warmhouse/apps/device-monitoring/api/telemetry-access"
	"github.com/voitenkov-courses/architecture-pro-warmhouse/apps/device-monitoring/internal/controllers/httpserver"
	telemetryprovider "github.com/voitenkov-courses/architecture-pro-warmhouse/apps/device-monitoring/internal/controllers/telemetry-provider"
	"github.com/voitenkov-courses/architecture-pro-warmhouse/apps/device-monitoring/internal/dal/queue"
	"github.com/voitenkov-courses/architecture-pro-warmhouse/apps/device-monitoring/internal/dal/storage"
)

var wg *sync.WaitGroup

func main() {
	// create a type that satisfies the `api.ServerInterface`, which contains an implementation of every operation from the generated code
	server := httpserver.NewServer()

	// Initialize router
	router := gin.Default()

	api.RegisterHandlers(router, server)

	srv := &http.Server{
		Addr:    getEnv("PORT", ":8080"),
		Handler: router,
	}

	storage := storage.New()

	storage.Connect()
	defer storage.Close()

	queue := queue.New()

	telemetryprovider := telemetryprovider.New(queue, storage)

	ctx, cancel := signal.NotifyContext(context.Background(),
		syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)
	defer cancel()

	go func() {
		<-ctx.Done()
		log.Println("Shutting down http server...")
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
		defer cancel()

		if err := srv.Shutdown(ctx); err != nil {
			log.Fatalf("Server forced to shutdown: %v\n", err)
		}

		log.Println("Server exited properly")
	}()

	wg = &sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()

		log.Printf("Server starting on %s\n", srv.Addr)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Failed to start HTTP Server: %v\n", err)
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()

		if err := telemetryprovider.Start(ctx); err != nil {
			log.Fatalf("Failed to start Telemetry Provider: %v\n", err)
		}
	}()

	wg.Wait()
	log.Println("Server exited properly")
	cancel()
	os.Exit(1) //nolint:gocritic
}

// getEnv gets an environment variable or returns a default value
func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}
