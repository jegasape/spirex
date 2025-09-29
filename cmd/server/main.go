package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	log.Println("Starting server ...")

	router := http.NewServeMux()

	response := struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
	}{
		Code:    http.StatusOK,
		Message: http.StatusText(http.StatusOK),
	}

	router.HandleFunc("/",
		func(w http.ResponseWriter, r *http.Request) {
			switch r.Method {
			case http.MethodGet:
				w.Header().Set("Content-type", "application/json")
				w.WriteHeader(http.StatusOK)
				defer w.(http.Flusher).Flush()
				if err := json.NewEncoder(w).Encode(&response); err != nil {
					http.Error(w, "Failed to encode", http.StatusInternalServerError)
					return
				}
			default:
				http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
				return
			}
		})

	server := http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutdown server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Fatal("Server down:", err.Error())
	}

	select {
	case <-ctx.Done():
		log.Println("Timeout of 5 second")
	case <-time.After(7 * time.Second):
		log.Println("Server exiting...")
	}
}
