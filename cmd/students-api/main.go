package main

import (
	"context"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/ageniouscoder/student-api/internal/config"
	"github.com/ageniouscoder/student-api/internal/http/handler/student"
)

func main() {
	//load config
	cfg := config.MustLoad()
	//database setup
	//setup route
	router := http.NewServeMux()
	router.HandleFunc("POST /api/students", student.New())
	//setup server
	server := http.Server{
		Addr:    cfg.Addr,
		Handler: router,
	}
	//fmt.Println("server started", cfg.Addr)
	slog.Info("server started at", slog.String("address", cfg.Addr))

	done := make(chan os.Signal, 1)

	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		err := server.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			log.Fatalf("failed to start server: %v", err)
		}
	}()

	<-done

	slog.Info("shutting down server gracefully")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	defer cancel()
	e := server.Shutdown(ctx)

	if e != nil {
		slog.Error("failed to shut down server", slog.String("error", e.Error()))
	}

	slog.Info("server shutdown Gracefully")

}
