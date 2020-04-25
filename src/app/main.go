package main

import (
	"log"
	"net/http"
	"os"
)

func main() {

	router := NewRouter()

	log.Fatal(http.ListenAndServe(os.Getenv("PORT"), router))
}
