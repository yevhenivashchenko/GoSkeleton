package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/joho/godotenv"
	"github.com/yevhenivashchenko/GoSkeleton/internal/config"
	"github.com/yevhenivashchenko/GoSkeleton/internal/handlers"
	"github.com/yevhenivashchenko/GoSkeleton/internal/server"
)

// @title GoSkeleton API
// @version 1.0
// @description Production-ready boilerplate for Go web services.
func main() {
	// 1. Load environment variables
	// Prioritizing .env file, fallback to OS environment if not present.
	if err := godotenv.Load(); err != nil {
		log.Printf("[INFO] .env file not found, defaulting to system env: %v", err)
	}

	// 2. Initialize application configuration
	// Config layer ensures all necessary environment variables are validated at startup.
	cfg := config.LoadConfig()

	// 3. Dependency Injection
	// Here we could inject database connections, loggers, or service layers.
	// For now, we instantiate the handler layer directly.
	h := handlers.NewHandler()

	// 4. Setup Router
	// Using standard library ServeMux, keeping dependencies minimal as per project philosophy.
	router := server.NewRouter(h)

	// 5. Configure HTTP Server
	srv := &http.Server{
		Addr:         ":" + cfg.Port,
		Handler:      router,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	// 6. Graceful Shutdown Implementation
	// Creating a channel to listen for OS interrupt signals (SIGINT, SIGTERM).
	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	// Run server in a goroutine so it doesn't block
	go func() {
		log.Printf("[INFO] Server is starting on port %s", cfg.Port)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("[FATAL] Could not listen on %s: %v\n", cfg.Port, err)
		}
	}()

	// Block until a signal is received
	<-done
	log.Println("[INFO] Server is stopping...")

	// Create a deadline to wait for current operations to complete
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("[FATAL] Server forced to shutdown: %v", err)
	}

	log.Println("[INFO] Server exited properly")
}
