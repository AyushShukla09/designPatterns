package main

import (
	"fmt"
	"log"
	"net/http"
)

const port = ":4000"

type Application struct{}

func main() {
	// app := Application{}
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Home Page")
	})
	fmt.Println("Starting server at port: ", port)
	if err := http.ListenAndServe(port, nil); err != nil {
		log.Println("Failed to start Server: ", err)
	}
}
