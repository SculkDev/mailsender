package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/getsentry/sentry-go"
)

func getEnv(key, fallback string) string {
	if val := os.Getenv(key); val != "" {
		return val
	}
	return fallback
}

func main() {
	dsn := os.Getenv("SENTRY_DSN")
	if dsn != "" {
		if err := sentry.Init(sentry.ClientOptions{
			Dsn: dsn,
		}); err != nil {
			log.Printf("sentry init error: %v", err)
		}
		defer sentry.Flush(2 * time.Second)
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/mailhook", mailhookHandler)

	addr := ":1432"
	log.Printf("starting mailsender on %s", addr)
	if err := http.ListenAndServe(addr, mux); err != nil {
		log.Fatalf("server error: %v", err)
	}
}
