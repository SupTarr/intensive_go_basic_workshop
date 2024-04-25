package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/SupTarr/intensive_go_basic_workshop/_exercises/thai_id"
)

type VerifyIdRequest struct {
	ID string `json:"id" binding:"required"`
}

type VerifyIdResponse struct {
	Valid bool `json:"valid"`
}

type VerifyIdErrorResponse struct {
	Message string `json:"message"`
}

func main() {
	slog.SetLogLoggerLevel(slog.LevelDebug)
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()
	idleConnsClosed := make(chan struct{})
	mux := http.NewServeMux()
	port := os.Getenv("PORT")

	srv := http.Server{
		Addr:              ":" + port,
		Handler:           mux,
		ReadHeaderTimeout: 5 * time.Second,
	}

	slog.Debug("Listening and serving HTTP on :" + port)
	go func() {
		<-ctx.Done()
		fmt.Println("Shutting down...")
		if err := srv.Shutdown(context.Background()); err != nil {
			log.Printf("HTTP server Shutdown: %v", err)
		}
		close(idleConnsClosed)
	}()
	if err := srv.ListenAndServe(); err != http.ErrServerClosed {
		log.Fatalf("HTTP server ListenAndServe: %v", err)
	}

	mux.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]string{
			"message": "pong",
		})
	})

	mux.HandleFunc("/thai/ids/verify", func(w http.ResponseWriter, r *http.Request) {
		var request VerifyIdRequest
		err := json.NewDecoder(r.Body).Decode(&request)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(VerifyIdErrorResponse{
				Message: err.Error(),
			})

			return
		}

		err = thai_id.Validate(request.ID)
		if err != nil {
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(VerifyIdResponse{
				Valid: false,
			})

			return
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(VerifyIdResponse{
			Valid: true,
		})
	})

	<-idleConnsClosed
	fmt.Println("Bye!")
}
