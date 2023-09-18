package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

type Config struct{}

func main() {
	app := &Config{}

	PORT := "8081"
	if port := os.Getenv("PORT"); port != "" {
		PORT = port
	}

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", PORT),
		Handler: app.routes(),
	}

	log.Println("Starting broker service on port", PORT)
	if err := srv.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
