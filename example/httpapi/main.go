package main

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"os"

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

	http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]string{
			"message": "pong",
		})
	})

	http.HandleFunc("/thai/ids/verify", func(w http.ResponseWriter, r *http.Request) {
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

	port := os.Getenv("PORT")
	slog.Debug("Listening and serving HTTP on :" + port)
	http.ListenAndServe(":"+port, nil)
}
