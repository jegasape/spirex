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
	log.Println("Starting server...")

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

			if r.URL.Path != "/" {
				http.NotFound(w, r)
				return
			}

			switch r.Method {
			case http.MethodGet:
				w.Header().Set("Content-type", "application/json")
				w.WriteHeader(http.StatusOK)

				if err := json.NewEncoder(w).Encode(&response); err != nil {
					http.Error(w, "Failed on response", http.StatusInternalServerError)
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
			log.Fatalf("Could not listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown -> ", err.Error())
	}

	<-ctx.Done()
	log.Println("Server exited gracefully")
}
