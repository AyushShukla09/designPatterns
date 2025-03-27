package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

const port = ":4000"

type Application struct{}

func main() {
	app := Application{}
	server := http.Server{
		Addr:         port,
		Handler:      app.routes(),
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
		IdleTimeout:  30 * time.Second,
	}
	fmt.Println("Starting server at port: ", port)
	if err := server.ListenAndServe(); err != nil {
		log.Println("Failed to start Server: ", err)
	}
}
