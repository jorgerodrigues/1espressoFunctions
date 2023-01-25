package main

import (
	"1espresso.com/initialRecipe"
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", initialrecipe.InitialRecipe)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Printf("Listening on port %s", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}
}
