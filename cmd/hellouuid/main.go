package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/robdimsdale/multimain"
)

func handler(w http.ResponseWriter, r *http.Request) {
	uuid := multimain.MustUUID()
	fmt.Fprintf(w, "Random UUID: %v", uuid)
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
