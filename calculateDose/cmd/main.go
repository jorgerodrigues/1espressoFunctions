package main

import (
	"log"
	"net/http"
	"os"

	"1espresso.com/dose"
)

func main() {
	http.HandleFunc("/", dose.Dose)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Printf("Listening on port %s", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}
}
